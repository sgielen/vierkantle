<template>
  <div class="cell" :class="{'unused' : props.state.used == 0}">
    <span class="letter">
      {{ props.letter }}
    </span>
    <span class="annotation firstLetter" v-if="props.state.begins && props.showBegins">
      {{ props.state.begins }}
    </span>
    <span class="annotation usedLetter" v-if="props.state.used && props.showUsed">
      {{ props.state.used }}
    </span>
  </div>
</template>

<script setup lang="ts">
import { CellState } from './models';

const props = defineProps<{
  letter: string,
  state: CellState,
  showBegins: boolean,
  showUsed: boolean,
}>();
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

  .letter {
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
</style>
