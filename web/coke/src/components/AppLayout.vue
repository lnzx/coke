<script setup>
import AppHeader from '@/components/AppHeader.vue'
</script>

<template>
  <div class="flex min-h-screen w-full flex-col">
    <AppHeader />

    <main class="flex min-h-[calc(100vh_-_theme(spacing.16))] flex-1 flex-col gap-4 p-4 md:gap-8 md:p-10">
      <slot></slot>
    </main>
  </div>
</template>
