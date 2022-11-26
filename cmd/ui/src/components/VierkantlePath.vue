<template>
  <canvas ref="canvas">
  </canvas>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted, ref, watchEffect } from 'vue';

const props = defineProps<{
  path: HTMLElement[],
}>();

const canvas = ref<HTMLCanvasElement>();
const canvasWidth = ref(canvas.value?.clientWidth ?? 0);
const canvasHeight = ref(canvas.value?.clientHeight ?? 0);
const canvasSizeInterval = ref();

onMounted(() => {
  canvasSizeInterval.value = setInterval(() => {
    const w = canvas.value?.clientWidth ?? 0;
    const h = canvas.value?.clientHeight ?? 0;
    if (w == 0 || h == 0) {
      // canvas is temporarily invisible, don't act on this
    } else if (canvasWidth.value != w || canvasHeight.value != h) {
      canvasWidth.value = w;
      canvasHeight.value = h;
      canvas.value!.width = w;
      canvas.value!.height = h;
    }
  }, 100);
});

onUnmounted(() => {
  clearInterval(canvasSizeInterval.value);
});

function middlePixelForElement(e: HTMLElement): {x: number; y: number} {
  const myRect = canvas.value?.getBoundingClientRect() ?? {x: 0, y: 0};
  const rect = e.getBoundingClientRect();
  return {
    x: (rect.x - myRect.x) + rect.width / 2,
    y: (rect.y - myRect.y) + rect.height / 2,
  };
}

watchEffect(() => {
  //emit('pathChanged', path.value);
  var ctx = canvas.value?.getContext('2d');
  if (!ctx) {
    return;
  }
  ctx.fillStyle = 'rgb(255, 0, 0)';
  ctx.strokeStyle = 'rgb(255, 0, 0)';
  ctx.lineWidth = 10;
  ctx.clearRect(0, 0, canvasWidth.value, canvasHeight.value);
  if (props.path.length == 0) {
    return;
  }

  const pixels = [] as {x: number, y: number}[];
  for (let i = 0; i < props.path.length; ++i) {
    const p = middlePixelForElement(props.path[i]);
    pixels.push(p);
    ctx.beginPath();
    ctx.ellipse(p.x, p.y, 20, 20, 0, 0, Math.PI * 2);
    ctx.fill();
  }

  ctx.moveTo(pixels[0].x, pixels[0].y);
  for (let i = 1; i < pixels.length; ++i) {
    ctx.lineTo(pixels[i].x, pixels[i].y);
  }
  ctx.stroke();
});
</script>

<style scoped lang="scss">
canvas {
  width: 100%;
  height: 100%;
}
</style>
