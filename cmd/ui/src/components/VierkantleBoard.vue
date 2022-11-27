<template>
  <div
    class="container"
    :style="css"
    @mouseup="dragEnd"
    @touchstart="touchEvent('start', $event)"
    @touchmove="touchEvent('move', $event)"
    @touchend="dragEnd"
  >
    <div class="board">
      <template v-for="row, y in props.board.cells">
        <template v-for="cell, x in row" :key="x + ',' + y">
          <VierkantleCell
            :letter="cell"
            :state="cellState(x, y)"
            :show-begins="true"
            :show-used="true"
            @mousedown="dragStart(x, y, $event.target)"
            @mouseenter="dragMove(x, y, $event.target)"
            @vk_touchstart="dragStart(x, y, $event.target)"
            @vk_touchmove="dragMove(x, y, $event.target)"
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
import { Board, CellState, Coord } from './models';
import VierkantleCell from './VierkantleCell.vue';
import VierkantlePath from './VierkantlePath.vue';

const props = withDefaults(defineProps<{
  board: Board,
}>(), {
});

const emit = defineEmits<{
  (event: 'word', word: string): void,
}>();

function cellState(x: number, y: number): CellState {
  let begins = 0;
  let used = 0;
  Object.values(props.board.words).forEach((word) => {
    if (!word.bonus && !word.guessed) {
      for (let i = 0; i < word.path.length; ++i) {
        if (word.path[i].x == x && word.path[i].y == y) {
          used += 1;
          if (i == 0) {
            begins += 1;
          }
        }
      }
    }
  });
  return { begins, used };
}

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

function dragStart(x: number, y: number, element: HTMLElement) {
  path.value = [{
    element,
    coord: {x, y},
  }];
}

function dragMove(x: number, y: number, element: HTMLElement) {
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

  // If the mouse is over a button that is reachable from the
  // last entry in path, add it
  const last = path.value[path.value.length-1];
  if (Math.abs(last.coord.x - x) <= 1 && Math.abs(last.coord.y - y) <= 1) {
    path.value.push({
      element,
      coord: { x, y },
    })
  }
}

function dragEnd() {
  const word = path.value.
    map((v) => props.board.cells[v.coord.y][v.coord.x]).
    reduce((prev, current) => prev + current, "");
  emit("word", word);
  path.value = [];
}

// While mouse drags can start on one element and then move to another element,
// triggering element1.mousedown() then element2.mouseenter(), touch drags keep
// triggering on element1 even if they move over element2. Therefore, we
// trigger touch events on .container, then move them downwards to the cells
// using a custom event. This custom event is ignored on all elements other
// than our VierkantleCells.
class VierkantleTouchStartEvent extends Event {}
class VierkantleTouchMoveEvent extends Event {}

function touchEvent(type: "start" | "move", event: TouchEvent) {
  var elements = document.elementsFromPoint(event.touches[0].pageX, event.touches[0].pageY)
  for(let element of elements) {
    if (type == "start") {
      element.dispatchEvent(new VierkantleTouchStartEvent("vk_touchstart"));
    } else {
      element.dispatchEvent(new VierkantleTouchMoveEvent("vk_touchmove"));
    }
  }
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
  touch-action: none;
}
</style>
