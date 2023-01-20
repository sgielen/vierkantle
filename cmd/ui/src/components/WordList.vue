<template>
  <template v-for="(words, length) in wordsByLength" :key="length">
    <template v-if="words && words.filter(f => props.configuring || !f[1].bonus).length">
      <h6>{{ length }} letters</h6>
      <template v-for="([word, wstate]) in words" :key="word">
        <template v-if="props.configuring">
          <div class="row items-center">
            <code>{{ word }}</code>
            <q-checkbox class="q-ml-sm" size="xs" label="Bonus" :model-value="wstate.bonus ?? false" @update:model-value="(value: boolean) => setBonus(word, value)" />
          </div>
        </template>
        <template v-else-if="!wstate.bonus">
          <template v-if="wstate.guessed">
            <code>{{ word }}</code><br/>
          </template>
          <template v-else>
            <code>{{ wordWithStars(word) }}</code><br />
          </template>
        </template>
      </template>
    </template>
  </template>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { WordInBoard } from './models';

const props = withDefaults(defineProps<{
  words: Record<string, WordInBoard>;
  configuring?: boolean,
}>(), {
  configuring: false,
})

const emit = defineEmits<{
  (event: 'set-bonus', word: string, bonus: boolean): void,
}>();

const wordsByLength = computed((): [string, WordInBoard][][] => {
  var res = [] as [string, WordInBoard][][]
  Object.entries(props.words).forEach(([word, wordInBoard]) => {
    res[word.length] ??= [];
    res[word.length].push([word, wordInBoard]);
  })
  for (let i = 0; i < res.length; ++i) {
    if (res[i] && res[i].length) {
      res[i] = res[i].sort((a, b) => { return a[0].length - b[0].length });
    }
  }
  return res;
});

function wordWithStars(w: string): string {
  function xStars(i: number): string {
    var s = ""
    for (; i > 0; --i) {
      s += "*"
    }
    return s
  }

  switch(w.length) {
    case 4: return w.substring(0, 1) + xStars(3)
    case 5: return w.substring(0, 1) + xStars(4)
    case 6: return w.substring(0, 2) + xStars(4)
    case 7: return w.substring(0, 2) + xStars(4) + w.substring(6)
    default:
      return w.substring(0, 2) + xStars(w.length - 4) + w.substring(w.length - 2)
  }
}

function setBonus(word: string, bonus: boolean) {
  emit("set-bonus", word, bonus);
}
</script>

<style scoped lang="scss">
h6 {
  margin: 0;
  margin-top: 20px;

  &:first-of-type {
    margin-top: 0;
  }
}
</style>
