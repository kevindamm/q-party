<template>
  <component :is="mediaType" href="mediaURL">{{ textContent }}</component>
</template>

<script setup lang="ts">
import type { Component } from "vue"
import { defineProps } from "vue"

defineProps<{
  textContent?: string,
  mediaType: Component,
  mediaURL: string,
}>()

</script>
