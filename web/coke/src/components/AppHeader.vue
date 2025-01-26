<script setup>
import { Input } from '@/components/ui/input'
import { CircleUser, Search, LogOut } from 'lucide-vue-next'
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu'
import AppNav from '@/components/AppNav.vue'
import { useUserSession } from '@/stores/userSession'

const userSession = useUserSession()

const logout = () => {
  userSession.clearToken()
  location.href = '/'
}
</script>

<template>
  <header class="sticky top-0 flex h-16 items-center gap-4 border-b bg-background px-4 md:px-6">
    <AppNav />

    <div class="flex w-full items-center gap-4 md:ml-auto md:gap-2 lg:gap-4">
      <form class="ml-auto flex-1 sm:flex-initial">
        <div class="relative">
          <Search class="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input type="search" placeholder="Search..." class="pl-8 sm:w-[300px] md:w-[200px] lg:w-[300px]" />
        </div>
      </form>
      <DropdownMenu>
        <DropdownMenuTrigger as-child>
          <Button variant="secondary" size="icon" class="rounded-full">
            <CircleUser class="h-5 w-5" />
            <span class="sr-only">Toggle user menu</span>
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end" class="w-40">
          <DropdownMenuLabel>My Account</DropdownMenuLabel>
          <DropdownMenuSeparator />
          <DropdownMenuItem @click="logout" class="flex justify-between items-center">
            <span>Log out</span>
            <DropdownMenuShortcut><LogOut class="h-5 w-5" /></DropdownMenuShortcut>
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  </header>
</template>
