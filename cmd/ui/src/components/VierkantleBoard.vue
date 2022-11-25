<template>
  <div class="container" :style="css" @mouseup="mouseup">
    <div class="board">
      <template v-for="row, y in props.board.cells">
        <template v-for="cell, x in row" :key="x + ',' + y">
          <VierkantleCell
            :letter="cell"
            @mouseover="mouseover(x, y, $event)"
            @mousedown="mousedown(x, y, $event)"
          />
        </template>
      </template>
    </div>

    <div class="path" style="pointer-events: none">
      <VierkantlePath
        :path="htmlPath"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { Board, Coord } from './models';
import VierkantleCell from './VierkantleCell.vue';
import VierkantlePath from './VierkantlePath.vue';

const props = withDefaults(defineProps<{
  board: Board,
}>(), {
});

const emit = defineEmits<{
  (event: 'word', word: string): void,
}>();

const css = computed(() => {
  return {
    '--board-width': props.board.width,
    '--board-height': props.board.height,
  }
})

interface PathEntry {
  element: HTMLElement;
  coord: Coord;
}

const path = ref<PathEntry[]>([]);

const htmlPath = computed(() => {
  return path.value.map((v) => v.element);
});

function mouseover(x: number, y: number, e: MouseEvent) {
  if (path.value.length == 0) {
    // not pathing, ignore mouseover
    return
  }

  // If the mouse is over the second-last entry in path, go back to it
  if (path.value.length >= 2) {
    const secondLast = path.value[path.value.length-2];
    if (secondLast.coord.x == x && secondLast.coord.y == y) {
      path.value.pop()
      return
    }
  }

  if (path.value.some((c) => c.coord.x == x && c.coord.y == y)) {
    // Already in path, don't add it
    return
  }

  // If path is empty, or the mouse is over an button that is reachable from the
  // last entry in path, add it
  const last = path.value.length == 0 ? undefined : path.value[path.value.length-1];
  if (last == undefined || Math.abs(last.coord.x - x) <= 1 && Math.abs(last.coord.y - y) <= 1) {
    path.value.push({
      element: e.target as HTMLElement,
      coord: { x, y },
    })
  }
}

function mousedown(x: number, y: number, e: MouseEvent) {
  path.value = [{
    element: e.target as HTMLElement,
    coord: {x, y},
  }]
}

function mouseup() {
  const word = path.value.
    map((v) => props.board.cells[v.coord.y][v.coord.x]).
    reduce((prev, current) => prev + current, "");
  emit("word", word);
  path.value = [];
}
</script>

<style scoped lang="scss">
/* The container decides the size of the board. It takes up
 * 100% of the available space, so that the caller determines
 * the size of this board.
 */
.container {
  position: relative;
  width: 100%;
  height: 100%;
}

/* .path displays the path overlay. */
.path {
  position: absolute;
  width: 100%;
  height: 100%;
  z-index: 9;
  opacity: 0.8;
}

/* .board displays the board. It takes up 100% of the available
 * space. The space inside is divided by a grid, with gaps
 * on all sides. */
.board {
  position: absolute;
  width: 100%;
  height: 100%;

  display: grid;

  --gap: 5%;
  column-gap: var(--gap);
  row-gap: var(--gap);

  --column-gap-size: calc(var(--gap) * (var(--board-height) - 1));
  --row-gap-size: calc(var(--gap) * (var(--board-width) - 1));

  grid-template-columns: repeat(var(--board-width), calc((100% - var(--column-gap-size)) / var(--board-width)));
  grid-template-rows: repeat(var(--board-height), calc((100% - var(--row-gap-size)) /var(--board-height)));

  justify-content: center;
  align-content: center;

  user-select: none;
  -moz-user-select: none;
  -khtml-user-select: none;
  -webkit-user-select: none;
  -o-user-select: none;

  cursor: default;
}
</style>
