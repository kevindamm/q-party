
<template>
  <table>
    <tbody>
      <tr v-for="player in players" :key="player">
    </tr></tbody>
  </table>
  <table>
    <tbody>
      <tr v-for="player in audience" :key="player">
    </tr></tbody>
  </table>
</template>

<script setup lang="ts">
import { ref } from 'vue';

// TODO 
const players = ref<string[] | undefined>();
const audience = ref<string[] | undefined>();


</script>


<style>
</style>