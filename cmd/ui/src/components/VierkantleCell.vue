<template>
  <div class="cell">
    <span class="letter">
      {{ props.letter }}
    </span>
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  letter: string,
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

  .letter {
    text-transform: uppercase;
  }
}
</style>
