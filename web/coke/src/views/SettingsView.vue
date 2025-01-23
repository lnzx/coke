<script setup lang="ts">
import { Button } from '@/components/ui/button'

import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card'
import { Checkbox } from '@/components/ui/checkbox'
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu'
import { Input } from '@/components/ui/input'
import { Sheet, SheetContent, SheetTrigger } from '@/components/ui/sheet'
import { Grip, CircleUser, Menu, Search, Server, Earth, Settings, Users } from 'lucide-vue-next'
</script>

<template>
  <div class="flex min-h-screen w-full flex-col">
    <header class="sticky top-0 flex h-16 items-center gap-4 border-b bg-background px-4 md:px-6">
      <nav class="hidden flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 md:text-sm lg:gap-6">
        <a href="#" class="flex items-center gap-2 text-lg font-semibold md:text-base">
          <Grip class="h-6 w-6" />
          <span class="sr-only">Acme Inc</span>
        </a>
        <a href="dash" class="flex items-center gap-1 text-muted-foreground transition-colors hover:text-foreground"> <Server class="h-4 w-4" /> Nodes </a>
        <a href="#" class="flex items-center gap-1 text-muted-foreground transition-colors hover:text-foreground"> <Users class="h-4 w-4" /> Users </a>
        <a href="#" class="flex items-center gap-1 text-muted-foreground transition-colors hover:text-foreground"> <Earth class="h-4 w-4" /> DNS </a>
        <a href="#" class="flex items-center gap-1 text-foreground transition-colors hover:text-foreground"> <Settings class="h-4 w-4" /> Settings </a>
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
            <a href="#" class="flex items-center gap-2 text-lg font-semibold">
              <Package2 class="h-6 w-6" />
              <span class="sr-only">Acme Inc</span>
            </a>
            <a href="dash" class="text-muted-foreground hover:text-foreground"> Nodes </a>
            <a href="#" class="text-muted-foreground hover:text-foreground"> Users </a>
            <a href="#" class="text-muted-foreground hover:text-foreground"> DNS </a>
            <a href="#" class="hover:text-foreground"> Settings </a>
          </nav>
        </SheetContent>
      </Sheet>
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
          <DropdownMenuContent align="end">
            <DropdownMenuLabel>My Account</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuItem>Settings</DropdownMenuItem>
            <DropdownMenuItem>Support</DropdownMenuItem>
            <DropdownMenuSeparator />
            <DropdownMenuItem>Logout</DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
    </header>
    <main class="flex min-h-[calc(100vh_-_theme(spacing.16))] flex-1 flex-col gap-4 p-4 md:gap-8 md:p-10">
      <div class="mx-auto grid w-full max-w-6xl gap-2">
        <h1 class="text-3xl font-semibold">Settings</h1>
      </div>
      <div class="mx-auto grid w-full max-w-6xl items-start gap-6 md:grid-cols-[180px_1fr] lg:grid-cols-[250px_1fr]">
        <nav class="grid gap-4 text-sm text-muted-foreground">
          <a href="#" class="font-semibold text-primary"> General </a>
          <a href="#"> Security </a>
          <a href="#"> Integrations </a>
          <a href="#"> Support </a>
          <a href="#"> Organizations </a>
          <a href="#"> Upgrade </a>
        </nav>
        <div class="grid gap-6">
          <Card>
            <CardHeader>
              <CardTitle>API KEY</CardTitle>
              <CardDescription> Used to identify your store in the marketplace. </CardDescription>
            </CardHeader>
            <CardContent>
              <form>
                <Input placeholder="api key" />
              </form>
            </CardContent>
            <CardFooter class="border-t px-6 py-4">
              <Button>Save</Button>
            </CardFooter>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Plugins Directory</CardTitle>
              <CardDescription> The directory within your project, in which your plugins are located. </CardDescription>
            </CardHeader>
            <CardContent>
              <form class="flex flex-col gap-4">
                <Input placeholder="Project Name" default-value="/content/plugins" />
                <div class="flex items-center space-x-2">
                  <Checkbox id="include" default-checked />
                  <label for="include" class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"> Allow administrators to change the directory. </label>
                </div>
              </form>
            </CardContent>
            <CardFooter class="border-t px-6 py-4">
              <Button>Save</Button>
            </CardFooter>
          </Card>
        </div>
      </div>
    </main>
  </div>
</template>
