<template>
  <div>hihihi</div>

  <p>
    let's get {{ episode }} started
  </p>

  <router-view />
</template>

<script setup lang="ts">
defineProps({episode: Number})
</script>