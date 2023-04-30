<template>
  <q-header elevated>
    <q-toolbar>
      <q-toolbar-title>
        Vierkantle Generator
      </q-toolbar-title>
    </q-toolbar>
  </q-header>

  <q-page-container>
    <q-page>
      <div class="game-page row items-center q-ma-xl">
        <div class="wordlist">
          <div class="row q-gutter-sm q-ma-sm">
            <q-input dense outlined type="number" label="Breedte" style="max-width: 80px;" v-model.number="width" />
            <q-input dense outlined type="number" label="Hoogte" style="max-width: 80px;" v-model.number="height" />
            <q-btn @click="reset" color="negative">Nieuw leeg bord</q-btn> <br />
            <q-input dense outlined type="text" label="Startwoord" style="max-width: 140px;" v-model="seedword" />
            <q-btn @click="seed" color="negative">Nieuw vooringevuld bord</q-btn> <br/>
          </div>
          <q-separator class="q-ma-sm" />
          <div class="row q-gutter-sm q-ma-sm">
            <q-btn @click="download" color="primary">Download bord</q-btn>
            <q-btn @click="chooseFile" color="primary">Upload bord</q-btn>
            <q-input filled type="file" ref="qFile" v-model="file" v-show="false" />
          </div>
          <q-separator class="q-ma-sm" />
          <div class="row q-gutter-sm q-ma-sm">
            <q-btn @click="fillIn" color="secondary">Vul bord verder in</q-btn>
            <q-btn @click="renewWords" color="secondary">Vernieuw woordenlijst</q-btn>
          </div>
          <q-separator class="q-ma-sm" />

          <div v-if="loading" class="q-my-md q-mx-sm">
            <q-linear-progress
              :indeterminate="!loadingProgress"
              :value="loadingProgress"
              rounded
              instant-feedback
              size="10px"
              color="primary"
            />
          </div>
          <WordList :words="board.words" :configuring="true" @set-bonus="setWordBonus" />
        </div>
        <div class="game-wrap items-center justify-evenly">
          <LabelAutofit size="48" :value="numWords" />
          <LabelAutofit size="24" :value="usage" />
          <LabelAutofit size="24" v-if="error" :value="error" />
          <div class="row items-center justify-evenly relative-position">
            <div class="game">
              <VierkantleBoard
                :board="board"
                @update:letter="updateLetter"
                :generatorMode="true"
              />
            </div>
            <q-inner-loading :showing="loading">
              <q-spinner-gears size="250px" color="primary" />
            </q-inner-loading>
          </div>
        </div>
      </div>
    </q-page>
  </q-page-container>
</template>

<script setup lang="ts">
import VierkantleBoard from 'components/VierkantleBoard.vue';
import { Board } from 'src/components/models';
import { onMounted, ref, computed } from 'vue';
import { StorageSerializers, useStorage } from '@vueuse/core';
import { createChannel, createClient } from 'nice-grpc-web';
import { VierkantleServiceDefinition, VierkantleServiceClient } from '../services/vierkantle';
import WordList from 'src/components/WordList.vue';
import LabelAutofit from 'src/components/LabelAutofit.vue';
import { QInput } from 'quasar';

function defaultBoard(): Board {
  const b: Board = {
    width: width.value || 4,
    height: height.value || 4,
    cells: [],
    words: {},
  };
  for (let y = 0; y < b.height; ++y) {
    b.cells[y] = [];
    for (let x = 0; x < b.width; ++x) {
      b.cells[y][x] = "";
    }
  }
  return b;
}

const width = ref(4);
const height = ref(4);
const board = useStorage<Board>("generatorBoard", defaultBoard(), undefined, { serializer: StorageSerializers.object });
onMounted(() => {
  width.value = board.value.width;
  height.value = board.value.height;
});

const numWords = computed(() => {
  const normalWords = Object.values(board.value.words).filter((w) => !w.bonus).length;
  const bonusWords = Object.values(board.value.words).filter((w) => w.bonus).length;
  return `${normalWords} woorden, ${bonusWords} bonuswoorden`;
});
const usage = computed(() => {
  return "Dubbelklik om een cel in te vullen. Cel leeg laten met <spatie>. Na wijzigen woordenlijst hernieuwen."
});

const error = ref("");
const loading = ref(false);
const loadingProgress = ref(0);
const backendAddress = window.location.origin + "/api";

function reset() {
  board.value = defaultBoard();
}

const seedword = ref("");

const boardFilename = computed(() => {
  if (seedword.value) {
    return seedword.value.trim().toLowerCase() + ".json";
  }
  return new Date().toISOString().slice(0, 10) + ".json";
})

async function seed() {
  loadingProgress.value = 0;
  loading.value = true;
  error.value = "";
  try {
    const channel = createChannel(backendAddress);
    const client: VierkantleServiceClient = createClient(VierkantleServiceDefinition, channel);
    const responses = client.seedBoard({
      seedWord: seedword.value,
      width: width.value,
      height: height.value,
      attempts: 10000,
    });
    for await (const response of responses) {
      if (response.board) {
        board.value = JSON.parse(new TextDecoder().decode(response.board));
      }
      if (response.progress) {
        loadingProgress.value = response.progress / response.attempts;
      }
    }
  } catch(e) {
    error.value = e as string;
  }
  loading.value = false;
}

async function fillIn() {
  loadingProgress.value = 0;
  loading.value = true;
  error.value = "";
  try {
    const channel = createChannel(backendAddress);
    const client: VierkantleServiceClient = createClient(VierkantleServiceDefinition, channel);
    const responses = client.fillInBoard({
      board: new TextEncoder().encode(JSON.stringify(board.value, null, 0)),
      attempts: 10000,
    });
    for await (const response of responses) {
      if (response.board) {
        board.value = JSON.parse(new TextDecoder().decode(response.board));
      }
      if (response.progress) {
        loadingProgress.value = response.progress / response.attempts;
      }
    }
  } catch(e) {
    error.value = e as string;
  }
  loading.value = false;
}

async function renewWords() {
  loading.value = true;
  error.value = "";
  try {
    const channel = createChannel(backendAddress);
    const client: VierkantleServiceClient = createClient(VierkantleServiceDefinition, channel);
    const boardResponse = await client.wordsForBoard({
      board: new TextEncoder().encode(JSON.stringify(board.value, null, 0)),
    });
    board.value = JSON.parse(new TextDecoder().decode(boardResponse.board));
  } catch(e) {
    error.value = e as string;
  }
  loading.value = false;
}

function updateLetter(x: number, y: number, letter: string) {
  board.value.cells[y][x] = letter;
}

function setWordBonus(word: string, bonus: boolean) {
  if (word in board.value.words) {
    board.value.words[word].bonus = bonus;
  }
}

function download() {
  const json = JSON.stringify(board.value, null, 0);
  const file = new File([json], boardFilename.value);

  const link = document.createElement("a");
  link.style.display = "none";
  link.href = URL.createObjectURL(file);
  link.download = file.name;

  document.body.appendChild(link);
  link.click();

  // To make this work on Firefox we need to wait a little while before removing
  // it.
  setTimeout(() => {
    URL.revokeObjectURL(link.href);
    link.parentNode?.removeChild(link);
  }, 0);
}

const qFile = ref<QInput>()
function chooseFile() {
  (qFile.value?.getNativeElement() as HTMLElement).click();
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const file = computed<any>({
  get: () => { return undefined },
  set: async (files: FileList | File | undefined) => {
    if (!files) {
      return;
    }
    const file: File = files.length ? (files as FileList)[0] : files as File;
    if (!('text' in file)) {
      return;
    }
    const json = await file.text();
    try {
      error.value = "";
      board.value = JSON.parse(json);
    } catch(e) {
      error.value = e as string;
    }
  },
});
</script>

<style lang="scss">
.game {
  min-width: 300px;
  width: 100%;
  max-width: min(600px, 80vw, 80vh);
  aspect-ratio: 1 / 1;

  margin: 20px 40px;
  font-size: 40px;
}

.wordlist {
  z-index: 500;
  background: white;
  border: 1px solid black;
  overflow: scroll;
  padding: 4px;
}

@media screen and (min-width: 500px) {
  .game-page {
    margin: 12px 48px;
  }

  .wordlist {
    width: 25%;
    height: auto;
    max-height: calc(100vh - 80px);
    display: block !important;
    visibility: visible !important;
  }

  .game-wrap {
    width: 75%;
    height: auto;
  }
}

@media screen and (max-height: 500px) {
  .game {
    font-size: 30px;
  }
}
</style>
