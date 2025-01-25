<script setup>
import { Button } from '@/components/ui/button'
import { Sheet, SheetContent, SheetTrigger } from '@/components/ui/sheet'
import { Grip, Package2, Server, Settings, Users, Menu } from 'lucide-vue-next'
import { RouterLink } from 'vue-router'
</script>

<template>
  <nav class="hidden flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 md:text-sm lg:gap-6">
    <RouterLink to="/nodes" class="flex items-center gap-2 text-lg font-semibold md:text-base">
      <Grip class="h-6 w-6" />
      <span class="sr-only">headscale</span>
    </RouterLink>

    <RouterLink to="/nodes" class="flex items-center gap-1 text-muted-foreground transition-colors hover:text-foreground"><Server class="h-4 w-4" /> Machines </RouterLink>
    <RouterLink to="/users" class="flex items-center gap-1 text-muted-foreground transition-colors hover:text-foreground"><Users class="h-4 w-4" /> Users </RouterLink>
    <RouterLink to="/settings" class="flex items-center text-muted-foreground gap-1 transition-colors hover:text-foreground"><Settings class="h-4 w-4" /> Settings </RouterLink>
  </nav>

  <Sheet>
    <SheetTrigger as-child>
      <Button variant="outline" size="icon" class="shrink-0 md:hidden">
        <Menu class="h-5 w-5" />
        <span class="sr-only">Toggle navigation menu</span>
      </Button>
    </SheetTrigger>
    <SheetContent side="left">
      <nav class="grid gap-6 text-lg font-medium">
        <RouterLink to="/nodes" class="flex items-center gap-2 text-lg font-semibold">
          <Package2 class="h-6 w-6" />
          <span class="sr-only">headscale</span>
        </RouterLink>
        <RouterLink to="/nodes" class="hover:text-foreground"> Nodes </RouterLink>
        <RouterLink to="/users" class="hover:text-foreground"> Users </RouterLink>
        <RouterLink to="/settings" class="hover:text-foreground"> Settings </RouterLink>
      </nav>
    </SheetContent>
  </Sheet>
</template>
