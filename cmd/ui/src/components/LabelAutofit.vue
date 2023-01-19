<template>
  <div ref="el">{{ value }}</div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';

const props = defineProps<{
  value: string;
  size: number | string;
}>();

const el = ref<HTMLDivElement>();

const fontSize = computed(() => {
  const spanStyle = window.getComputedStyle(el.value!, null);
  const fontWeight = spanStyle.getPropertyValue('font-weight') || 'normal';
  const fontFamily = spanStyle.getPropertyValue('font-family') || 'sans-serif';
  const canvas = document.createElement("canvas");
  const ctx = canvas.getContext("2d");

  let fontSize = props.size;
  if (typeof fontSize === "string") {
    fontSize = parseInt(fontSize, 10);
  }
  if (!ctx) {
    return fontSize;
  }

  while (fontSize > 0) {
    ctx.font = `${fontWeight} ${fontSize}px ${fontFamily}`;
    const metrics = ctx.measureText(props.value);
    if (metrics.width <= window.innerWidth) {
      return `${fontSize}px`;
    }
    fontSize -= 1;
  }
  return 1;
});
</script>

<style scoped lang="scss">
div {
  margin: 0;
  padding: 0;
  font-size: v-bind(fontSize);
  line-height: v-bind(fontSize);
  text-align: center;
  white-space: nowrap;
}
</style>
