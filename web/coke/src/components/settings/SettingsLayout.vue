<script setup>
import { RouterLink } from 'vue-router'
</script>

<template>
  <div class="mx-auto grid w-full max-w-6xl gap-2">
    <h1 class="text-3xl font-semibold">Settings</h1>
  </div>
  <div class="mx-auto grid w-full max-w-6xl items-start gap-6 md:grid-cols-[180px_1fr] lg:grid-cols-[250px_1fr]">
    <nav class="grid gap-4 text-sm text-muted-foreground">
      <RouterLink to="/settings" class="hover:text-foreground"> General </RouterLink>
      <RouterLink to="/settings/keys" class="hover:text-foreground"> Keys </RouterLink>
      <RouterLink to="/settings/node" class="hover:text-foreground"> Device management </RouterLink>
      <RouterLink to="/settings/dns" class="hover:text-foreground"> DNS </RouterLink>
    </nav>
    <div class="grid gap-6">
      <slot></slot>
    </div>
  </div>
</template>
