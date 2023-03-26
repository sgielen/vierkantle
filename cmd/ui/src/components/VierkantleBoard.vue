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
            :generator-mode="props.generatorMode"
            @mousedown="dragStart(x, y, $event)"
            @mousemove="dragMove(x, y, $event)"
            @vk_touchstart="dragStart(x, y, $event)"
            @vk_touchmove="dragMove(x, y, $event)"
            @update:letter="updateLetter(x, y, $event)"
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
  generatorMode?: boolean,
}>(), {
  generatorMode: false,
});

const emit = defineEmits<{
  // A word was guessed. Only when generatorMode is false.
  (event: 'word', word: string): void,

  // A word is being guessed. Only when generatorMode is false.
  (event: 'partialWord', word: string): void,

  // Emit an updated letter in the board. Only when generatorMode is true.
  (event: 'update:letter', x: number, y: number, value: string): void,
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

const wordPath = computed(() => {
  return path.value.
    map((v) => props.board.cells[v.coord.y][v.coord.x]).
    reduce((prev, current) => prev + current, "");
})

function dragStart(x: number, y: number, event: MouseEvent | VierkantleTouchEvent) {
  if (!props.generatorMode && props.board.cells[y][x] === "") {
    // skip empty cells in normal mode
    return
  }

  path.value = [{
    element: event.currentTarget as HTMLElement,
    coord: {x, y},
  }];
  emit("partialWord", wordPath.value);
}

function dragMove(x: number, y: number, event: MouseEvent | VierkantleTouchEvent) {
  if (path.value.length == 0) {
    // not pathing, ignore mouseover
    return
  }

  if (props.generatorMode) {
    // don't path in generator mode
    return;
  }

  if (props.board.cells[y][x] === "") {
    // skip empty cells
    return
  }

  // Only trigger on moves that are close to the center of the element
  const element = event.currentTarget as HTMLElement;
  const elementRect = element.getBoundingClientRect();
  const xFactor = (event.clientX - elementRect.x) / elementRect.width;
  const yFactor = (event.clientY - elementRect.y) / elementRect.height;
  // The distance from the center of the element
  const distFromCenter = Math.sqrt(Math.pow(xFactor - 0.5, 2) + Math.pow(yFactor - 0.5, 2));
  if (distFromCenter > 0.35) {
    // Event occurred outside of the action zone, so ignore it
    return
  }

  // If the mouse is over the second-last entry in path, go back to it
  if (path.value.length >= 2) {
    const secondLast = path.value[path.value.length-2];
    if (secondLast.coord.x == x && secondLast.coord.y == y) {
      path.value.pop()
      emit("partialWord", wordPath.value);
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
    emit("partialWord", wordPath.value);
  }
}

function dragEnd() {
  emit("word", wordPath.value);
  path.value = [];
}

// While mouse drags can start on one element and then move to another element,
// triggering element1.mousedown() then element2.mouseenter(), touch drags keep
// triggering on element1 even if they move over element2. Therefore, we
// trigger touch events on .container, then move them downwards to the cells
// using a custom event. This custom event is ignored on all elements other
// than our VierkantleCells.
class VierkantleTouchEvent extends Event {
  public clientX!: number;
  public clientY!: number;
}

function touchEvent(type: "start" | "move", event: TouchEvent) {
  event.preventDefault();
  var elements = document.elementsFromPoint(event.touches[0].clientX, event.touches[0].clientY)
  for(let element of elements) {
    const vkEvent = new VierkantleTouchEvent("vk_touch" + type);
    vkEvent.clientX = event.touches[0].clientX;
    vkEvent.clientY = event.touches[0].clientY;
    element.dispatchEvent(vkEvent);
  }
}

function updateLetter(x: number, y: number, letter: string) {
  if (props.generatorMode) {
    emit("update:letter", x, y, letter);
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
