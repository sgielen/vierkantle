<template>
  <q-page class="row items-center justify-evenly">
    <VierkantleBoard
      :board="board"
      @word="word"
    />
  </q-page>
</template>

<script setup lang="ts">
import VierkantleBoard from 'components/VierkantleBoard.vue';
import { Board } from 'src/components/models';
import { ref } from 'vue';

const board = ref<Board>({
  width: 4,
  height: 4,
  cells: [
    ['o', 'l', 'p', 'z', ],
    ['a', 'g', 'o', 't', ],
    ['s', 'r', 'i', 's', ],
    ['a', 't', 'm', 'e', ],
  ],
});

function word(word: string) {
  alert(word);
}
</script>
