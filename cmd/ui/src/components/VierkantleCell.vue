<template>
  <div class="cell">
    {{ props.letter }}
  </div>
</template>

<script setup lang="ts">
const props = defineProps<{
  letter: string,
}>();
</script>

<style lang="scss">
.cell {
  line-height: 3em;
  font-size: 32pt;
  text-transform: uppercase;
  border: 4px solid black;
  border-radius: 5%;
  text-align: center;
  vertical-align: middle;
}
</style>
