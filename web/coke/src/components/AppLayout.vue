<script setup lang="ts">
import AppHeader from './AppHeader.vue'
</script>

<template>
  <div class="flex min-h-screen w-full flex-col">
    <AppHeader />

    <slot></slot>
  </div>
</template>
