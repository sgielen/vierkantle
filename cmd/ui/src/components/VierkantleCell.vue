<template>
  <div
    class="cell"
    :class="{'unused': props.state.used == 0, 'empty': props.letter == '', 'generator': props.generatorMode}"
    @dblclick="doubleClicked"
    v-touch-hold="doubleClicked"
  >
    <span v-if="!editing" class="letter">
      {{ props.letter }}
    </span>
    <input v-if="editing" ref="letterInput" class="letter-input" type="text" :placeholder="props.letter" @keyup="keyUp" @blur="blur" />

    <span class="annotation firstLetter" v-if="props.state.begins && props.showBegins">
      {{ props.state.begins }}
    </span>
    <span class="annotation usedLetter" v-if="props.state.used && props.showUsed">
      {{ props.state.used }}
    </span>
  </div>
</template>

<script setup lang="ts">
import { ref, nextTick } from 'vue';
import { CellState } from './models';

const props = withDefaults(defineProps<{
  letter: string,
  state: CellState,
  showBegins: boolean,
  showUsed: boolean,
  generatorMode?: boolean,
}>(), {
  generatorMode: false,
});

const emit = defineEmits<{
  (event: 'update:letter', letter: string): void,
}>();

const editing = ref(false);
const letterInput = ref<HTMLInputElement>();

function doubleClicked() {
  if (props.generatorMode) {
    editing.value = true;
    nextTick(() => letterInput.value?.focus());
  }
}

function keyUp(event: KeyboardEvent) {
  if (props.generatorMode) {
    if (event.key === 'Shift' || event.key === 'Control' || event.key === 'Alt') {
      // don't stop editing when pressing modifiers
      return;
    }
    editing.value = false;
    if (event.key === ' ') {
      emit("update:letter", "");
    } else if (event.key.length === 1) {
      emit("update:letter", event.key.toLowerCase());
    }
  }
}

function blur() {
  editing.value = false;
}
</script>

<style scoped lang="scss">
/* A cell takes up all its space, then renders the letter
 * inside according to the font-size, which is taken from
 * the parent element of the VierkantleBoard. */
.cell {
  position: relative;
  width: 100%;
  height: 100%;
  border: 4px solid black;
  border-radius: 5%;
  display: flex;
  justify-content: center;
  align-items: center;

  &.unused {
    background-color: gray;
  }

  &.empty {
    border: none;
  }

  &.empty.unused {
    background-color: transparent !important;
  }

  &.empty.generator {
    border: 1px dashed black !important;
  }

  .letter {
    text-transform: uppercase;
  }

  .letter-input {
    width: 1em;
    text-transform: uppercase;
  }

  .annotation {
    position: absolute;
    font-size: 50%;
    line-height: 1;
    bottom: 0;
  }

  .firstLetter {
    left: 5%;
    color: red;
  }

  .usedLetter {
    right: 5%;
  }
}

body.body--dark .cell {
  border: 2px solid white;

  &.unused {
    background-color: #545454;
  }

  &.empty.unused {
    background-color: transparent !important;
  }

  &.empty.generator {
    border: 1px dashed white !important;
  }

  .firstLetter {
    color: #fffe00;
  }

}
</style>
