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
          <q-btn @click="reset" color="negative">Verwijder bord en begin opnieuw</q-btn> <br />
          <q-btn @click="renewWords" color="primary">Vernieuw woordenlijst</q-btn> <br />
          <q-btn @click="download" color="primary">Download bord</q-btn> <br />
          <template v-if="loading">
            <q-circular-progress
              indeterminate
              rounded
              size="50px"
              color="primary"
              class="q-ma-md"
            />
          </template>
          <WordList v-else :words="board.words" :showAll="true" />
        </div>
        <div class="game-wrap items-center justify-evenly">
          <!-- <VierkantleTop v-if="board" :board="board" :partialWord="partialWord" :lastWords="lastWords" /> -->
          <div class="row items-center justify-evenly">
            <div v-if="error">{{ error }}</div>
            <div class="game">
              <VierkantleBoard
                :board="board"
                @update:letter="updateLetter"
                :generatorMode="true"
              />
            </div>
          </div>
        </div>
      </div>
    </q-page>
  </q-page-container>
</template>

<script setup lang="ts">
import VierkantleBoard from 'components/VierkantleBoard.vue';
import { Board } from 'src/components/models';
import { ref } from 'vue';
import { StorageSerializers, useStorage } from '@vueuse/core';
import { createChannel, createClient } from 'nice-grpc-web';
import { VierkantleServiceDefinition, VierkantleServiceClient } from '../services/vierkantle';
import WordList from 'src/components/WordList.vue';

const defaultBoard: Board = {
  width: 4,
  height: 4,
  cells: [
    ["", "", "", ""],
    ["", "", "", ""],
    ["", "", "", ""],
    ["", "", "", ""],
  ],
  words: {},
};

const board = useStorage<Board>("generatorBoard", defaultBoard, undefined, { serializer: StorageSerializers.object });

const error = ref("");
const loading = ref(false);
const backendAddress = window.location.origin + "/api";

function reset() {
  board.value = defaultBoard;
}

async function renewWords() {
  loading.value = true;
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

function download() {
  const json = JSON.stringify(board.value, null, 0);
  const file = new File([json], "board.json");

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
