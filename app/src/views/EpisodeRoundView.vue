<template>
  <div>hi ({{ round }} jeopardy)</div>
</template>

<script setup lang="ts">
defineProps({round: String})
</script>