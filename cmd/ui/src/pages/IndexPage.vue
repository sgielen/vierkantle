<template>
  <q-page>
    <div class="row items-center justify-evenly">{{ wordsRemaining }} words remaining</div>
    <div class="row items-center justify-evenly">{{ wordMessage }}</div>
    <div class="row items-center justify-evenly">
      <div v-if="error">{{ error }}</div>
      <div v-else-if="!board">Loading board...</div>
      <VierkantleBoard
        v-else
        :board="board"
        @word="word"
      />
    </div>
  </q-page>
</template>

<script setup lang="ts">
import VierkantleBoard from 'components/VierkantleBoard.vue';
import { Board } from 'src/components/models';
import { computed, onMounted, ref } from 'vue';
import { StorageSerializers, useStorage } from '@vueuse/core';

const board_ = useStorage<Board | undefined>("board", undefined, undefined, { serializer: StorageSerializers.object });

const board = computed(() => {
  return board_.value;
});

const error = ref("");
onMounted(async () => {
  const today = new Date().toISOString().slice(0, 10);

  if (board_.value && board_.value.loadedAt === today) {
    // Today's board is already present, keep it
    return
  }

  // Download a new board
  board_.value = null;
  try {
    const f = await fetch("/board.json");
    board_.value = await f.json() as Board;
    board_.value.loadedAt = today;
  } catch(e) {
    error.value = e as string;
  }
});

const wordMessage = ref(".");

const wordsRemaining = computed(() => {
  if (!board.value) {
    return 0;
  }
  let r = 0;
  Object.values(board.value.words).forEach((word) => {
    if (!word.bonus && !word.guessed) {
      r += 1;
    }
  });
  return r;
});

function word(word: string) {
  if (word.length < 4) {
    wordMessage.value = "te kort: " + word
    return
  }
  const wordInBoard = board.value!.words[word];
  if (!wordInBoard) {
    wordMessage.value = "nope: " + word
    return
  }
  if (wordInBoard.bonus) {
    wordMessage.value = "bonus: " + word
  } else {
    wordMessage.value = word
  }
  if (wordInBoard.guessed) {
    wordMessage.value += " (had je al)"
  }
  wordInBoard.guessed = true
}
</script>
