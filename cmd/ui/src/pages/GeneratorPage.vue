<template>
  <q-header elevated>
    <q-toolbar>
      <q-toolbar-title>
        Vierkantle Generator
      </q-toolbar-title>

      <q-toggle v-model="dark" icon="dark_mode" color="secondary" />
    </q-toolbar>
  </q-header>

  <q-page-container>
    <q-page>
      <div class="game-page row items-center q-ma-xl">
        <div class="wordlist">
          <div class="row q-gutter-sm q-ma-sm">
            <div class="col-12">
              <p>Dit bord is gemaakt door (optioneel): <q-input dense outlined v-model="board.madeBy" /></p>
              <p>Beschrijving (optioneel):</p>
              <textarea v-model="board.description" style="width: 100%; height: 4em;" />
            </div>
          </div>
          <q-separator class="q-ma-sm" />
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
            <q-btn @click="addToQueue" color="primary">Voeg bord aan wachtrij toe</q-btn>
            <q-btn @click="chooseFile" color="primary">Upload bord</q-btn>
            <q-input filled type="file" ref="qFile" v-model="file" v-show="false" />
            <q-btn @click="queueOpen = true" color="primary" v-if="whoami?.admin">Open wachtrij</q-btn>
            <p v-if="boardStats">
              Laatste bord op dit moment: {{ boardStats.lastBoardName }}<br/>
              Borden in wachtrij: {{ boardStats.boardsInQueue }}
            </p>
          </div>
          <q-separator class="q-ma-sm" />
          <div class="row q-gutter-sm q-ma-sm">
            <q-btn @click="fillIn" color="secondary">Vul bord verder in</q-btn>
            <q-btn @click="renewWords" color="secondary">Vernieuw woordenlijst</q-btn>
          </div>
          <q-separator class="q-ma-sm" />

          <div class="q-ma-xs q-pa-xs" style="border: 1px solid rgb(255, 155, 111);" v-if="!username">
            <strong>Let op: je bent niet ingelogd.</strong> Als je ingelogd
            bent, dan onthoudt de generator alle woordtypes van woorden die je
            hier aan- of uitvinkt, maar momenteel dus niet. Je kan inloggen op
            de <a href="/">hoofdpagina</a> via het scorebord. Geen zorgen, de
            generator onthoudt ondertussen waar je mee bezig was.
          </div>

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
          <LabelAutofit size="24" class="error" v-if="error" :value="error" />
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

  <q-dialog v-model="queuedOpen">
    <q-card style="width: 650px">
      <q-card-section>
        <span class="text-h6">Bord toegevoegd aan wachtrij!</span>
      </q-card-section>
      <q-separator />
      <q-card-section class="q-pa-md">
        <template v-if="error">
          <span style="color: red">{{ error }}</span>
        </template>
        <template v-else>
          Je bord is toegevoegd aan de wachtrij en heeft id {{ queueId }} gekregen. Daarnaast heeft Sjors een mailtje gekregen.
          Bedankt voor het opsturen!
        </template>
      </q-card-section>
    </q-card>
  </q-dialog>

  <q-dialog v-model="queueOpen">
    <BoardQueue
      :backend="backendAddress"
      :currentBoard="boardJson"
      @open="openFromQueue"
    />
  </q-dialog>

</template>

<script setup lang="ts">
import VierkantleBoard from 'components/VierkantleBoard.vue';
import { Board } from 'src/components/models';
import { onMounted, ref, computed, watch } from 'vue';
import { StorageSerializers, useStorage } from '@vueuse/core';
import { createChannel, createClient } from 'nice-grpc-web';
import { VierkantleServiceDefinition, VierkantleServiceClient, GetBoardFromQueueResponse, GetNewestBoardResponse } from '../services/vierkantle';
import WordList from 'src/components/WordList.vue';
import LabelAutofit from 'src/components/LabelAutofit.vue';
import { QInput } from 'quasar';
import { isAbortError, useUniqueCall } from 'src/services/abort';
import { errorToString } from 'src/services/errors';
import { useQuasar } from 'quasar';
import { useWhoami } from 'src/services/whoami';
import BoardQueue from 'src/components/BoardQueue.vue';

const $q = useQuasar();
const dark = computed({
  get() {
    return $q.dark.isActive;
  },
  set(status: boolean) {
    $q.dark.set(status);
  },
});

const backendAddress = window.location.origin + "/api";
const channel = createChannel(backendAddress);
const client: VierkantleServiceClient = createClient(VierkantleServiceDefinition, channel);
const { username, whoami, updateWhoami } = useWhoami(client);
const myName = useStorage<string | undefined>("generatorName", undefined);

onMounted(async () => {
  await updateWhoami();
  if (!myName.value && username.value) {
    myName.value = username.value;
  }

  await updateStats();
});

const boardStats = ref<GetNewestBoardResponse>();
async function updateStats() {
  try {
    boardStats.value = await client.getNewestBoard({});
  } catch(e) {
    error.value = errorToString(e);
  }
}

function defaultBoard(): Board {
  const b: Board = {
    width: width.value || 4,
    height: height.value || 4,
    cells: [],
    words: {},
    madeBy: myName.value,
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

watch(() => board.value.madeBy, () => {
  if (board.value.madeBy) {
    myName.value = board.value.madeBy;
  }
});

const numWords = computed(() => {
  const normalWords = Object.values(board.value.words).filter((w) => !w.bonus).length;
  const bonusWords = Object.values(board.value.words).filter((w) => w.bonus).length;
  return `${normalWords} woorden, ${bonusWords} bonuswoorden`;
});
const usage = computed(() => {
  return "Dubbelklik om een cel in te vullen. Cel leeg maken met <spatie>. Na wijzigen woordenlijst hernieuwen."
});

const error = ref("");
const loading = ref(false);
const loadingProgress = ref(0);

const { newUniqueCall } = useUniqueCall();

function reset() {
  board.value = defaultBoard();
}

const seedword = ref("");

const boardName = computed(() => {
  if (seedword.value) {
    return seedword.value.trim().toLowerCase();
  }
  return new Date().toISOString().slice(0, 10);
})

const boardFilename = computed(() => {
  return boardName.value + ".json";
})

function replaceBoardFromApi(newBoard: Uint8Array) {
  const oldBoard = board.value;
  board.value = JSON.parse(new TextDecoder().decode(newBoard));
  board.value.madeBy = oldBoard.madeBy;
  board.value.description = oldBoard.description;
}

async function seed() {
  loadingProgress.value = 0;
  loading.value = true;
  error.value = "";
  try {
    const responses = client.seedBoard({
      seedWord: seedword.value,
      width: width.value,
      height: height.value,
      attempts: 10000,
    }, { signal: newUniqueCall() });
    for await (const response of responses) {
      if (response.board) {
        replaceBoardFromApi(response.board);
      }
      if (response.progress) {
        loadingProgress.value = response.progress / response.attempts;
      }
    }
  } catch(e) {
    if (isAbortError(e)) {
      return;
    }

    error.value = errorToString(e);
  }
  loading.value = false;
}

async function fillIn() {
  loadingProgress.value = 0;
  loading.value = true;
  error.value = "";
  try {
    const responses = client.fillInBoard({
      board: boardJson.value,
      attempts: 10000,
    }, { signal: newUniqueCall() });
    for await (const response of responses) {
      if (response.board) {
        replaceBoardFromApi(response.board);
      }
      if (response.progress) {
        loadingProgress.value = response.progress / response.attempts;
      }
    }
  } catch(e) {
    if (isAbortError(e)) {
      return;
    }

    error.value = errorToString(e);
  }
  loading.value = false;
}

async function renewWords() {
  loading.value = true;
  error.value = "";
  try {
    const boardResponse = await client.wordsForBoard({
      board: boardJson.value,
    }, { signal: newUniqueCall() });
    replaceBoardFromApi(boardResponse.board);
  } catch(e) {
    if (isAbortError(e)) {
      return;
    }

    error.value = errorToString(e);
  }
  loading.value = false;
}

function updateLetter(x: number, y: number, letter: string) {
  board.value.cells[y][x] = letter;
}

async function setWordBonus(word: string, bonus: boolean) {
  if (word in board.value.words) {
    board.value.words[word].bonus = bonus;
    try {
      await client.markWordType({ word, bonus });
    } catch(e) {
      console.log("markWordType error: ", e);
    }
  }
}

const boardJson = ref(new Uint8Array);
watch(board, () => {
  boardJson.value = new TextEncoder().encode(JSON.stringify(board.value, null, 0));
}, { deep: true });

function download() {
  const file = new File([boardJson.value], boardFilename.value);

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

const queuedOpen = ref(false);
const queueId = ref<number>();

async function addToQueue() {
  try {
    error.value = "";
    const response = await client.addBoardToQueue({
      boardName: boardName.value,
      board: boardJson.value,
    })
    queueId.value = response.id;
    queuedOpen.value = true;
    await updateStats();
  } catch(e) {
    error.value = errorToString(e);
  }
}

const queueOpen = ref(false);

function openFromQueue(newBoard: GetBoardFromQueueResponse) {
  board.value = JSON.parse(new TextDecoder().decode(newBoard.board));
  queueOpen.value = false;
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
      error.value = errorToString(e);
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

.error {
  color: red;
}

body.body--dark {
  header {
    background-color: #0d4174;
  }

  .wordlist {
    background: var(--q-dark-page);
    border: 1px solid white;
  }
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
