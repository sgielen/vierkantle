<template>
  <div class="outer">
    <div class="lastwords">
      <TransitionGroup name="fade" tag="div">
        <div v-for="(lastWord, i) in lastNWords" :key="i" class="lastword">
          <span>{{ lastWord }}</span>
        </div>
      </TransitionGroup>
    </div>
    <LabelAutofit class="progress" size="48" :value="progress" />
    <LabelAutofit class="message" size="34" :value="message" />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import LabelAutofit from './LabelAutofit.vue';
import { Board } from './models';

const props = defineProps<{
  board: Board;
  partialWord: string;
  lastWords: string[];
}>();

const wordsTotal = computed(() => {
  return Object.values(props.board.words).filter((w) => !w.bonus).length;
});

const wordsGuessed = computed(() => {
  if (!props.board) {
    return 0;
  }

  return Object.values(props.board.words).filter((w) => !w.bonus && w.guessed).length;
});

const progress = computed(() => {
  if (wordsGuessed.value == wordsTotal.value) {
    return "Klaar!";
  } else {
    return `${wordsGuessed.value} van ${wordsTotal.value}`;
  }
});

const message = computed(() => {
  return props.partialWord;
});

const lastNWords = computed(() => {
  let res = [] as string[];
  for (let i = props.lastWords.length - 1; i >= 0; --i) {
    res.push(props.lastWords[i]);
    if (res.length > 3) {
      return res;
    }
  }
  return res;
});
</script>

<style scoped lang="scss">
.progress {
  margin-top: 20px;
  height: 3rem;
}
.message {
  margin-top: 10px;
  height: 2.5rem;
}

.outer {
  position: relative;

  .lastwords {
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;

    .lastword {
      margin-bottom: 20px;
      text-align: center;

      span {
        background-color: white;
        border: 3px solid var(--q-primary);
        padding: 8px;
        font-size: 1.5em;
      }

      &.fade-enter-active,
      &.fade-leave-active {
        transition: opacity 0.5s ease;
      }

      &.fade-enter-from,
      &.fade-leave-to {
        opacity: 0;
      }
    }
  }
}
</style>
