<script setup lang="ts">
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu'
import { Input } from '@/components/ui/input'
import { Sheet, SheetContent, SheetTrigger } from '@/components/ui/sheet'
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog'
import { Plus, CircleUser, Menu, Grip, Search, Server, Earth, Settings, Ellipsis, Users, CopyIcon } from 'lucide-vue-next'
import { ref, onMounted } from 'vue'

const logout = () => {
  location.href = '/'
}

interface Node {
  id: number
  name: string
  addr: string
  status: boolean
}

const nodes = ref<Node[]>([])

const fetchNodes = async () => {
  const response = await fetch('https://676e290ddf5d7dac1cc9d14a.mockapi.io/api/nodes')
  const data: Node[] = await response.json()
  nodes.value = data
}

onMounted(() => {
  fetchNodes()
})
</script>

<template>
  <div class="flex min-h-screen w-full flex-col">
    <header class="sticky top-0 flex h-16 items-center gap-4 border-b bg-background px-4 md:px-6">
      <nav class="hidden flex-col gap-6 text-lg font-medium md:flex md:flex-row md:items-center md:gap-5 md:text-sm lg:gap-6">
        <a href="#" class="flex items-center gap-2 text-lg font-semibold md:text-base">
          <Grip class="h-6 w-6" />
          <span class="sr-only">headscale</span>
        </a>

        <a href="dash" class="flex items-center gap-1 transition-colors hover:text-foreground"><Server class="h-4 w-4" /> Nodes </a>
        <a href="users" class="flex items-center gap-1 text-muted-foreground transition-colors hover:text-foreground"><Users class="h-4 w-4" /> Users </a>
        <a href="dns" class="flex items-center gap-1 text-muted-foreground transition-colors hover:text-foreground"><Earth class="h-4 w-4" /> DNS </a>
        <a href="settings" class="flex items-center text-muted-foreground gap-1 transition-colors hover:text-foreground"><Settings class="h-4 w-4" /> Settings </a>
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
              <span class="sr-only">headscale</span>
            </a>
            <a href="dash" class="hover:text-foreground"> Nodes </a>
            <a href="#" class="text-muted-foreground hover:text-foreground"> Users </a>
            <a href="#" class="text-muted-foreground hover:text-foreground"> DNS </a>
            <a href="#" class="text-muted-foreground hover:text-foreground"> Settings </a>
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
            <DropdownMenuItem>Profile</DropdownMenuItem>
            <DropdownMenuSeparator />
            <DropdownMenuItem @click="logout">Logout</DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
    </header>

    <main class="flex flex-1 flex-col gap-4 p-4 md:gap-8 md:p-8">
      <div class="grid gap-4 md:gap-8">
        <Card class="xl:col-span-2">
          <CardHeader class="flex flex-row items-center">
            <div class="grid gap-2">
              <CardTitle class="text-2xl font-semibold">Nodes</CardTitle>
              <CardDescription>
                <Badge variant="secondary">{{ nodes.length }} nodes</Badge>
              </CardDescription>
            </div>
            <div class="ml-auto gap-1">
              <Dialog class="ml-auto gap-1">
                <DialogTrigger as-child>
                  <Button size="sm"> Add node <Plus class="h-4 w-4" /> </Button>
                </DialogTrigger>
                <DialogContent class="sm:max-w-xl">
                  <DialogHeader>
                    <DialogTitle>Add node</DialogTitle>
                    <DialogDescription> To add a new device to your Tailnet, simply install Tailscale on that device and run the following command. </DialogDescription>
                  </DialogHeader>
                  <div class="flex items-center space-x-2">
                    <div class="grid flex-1 gap-2">
                      <Label for="link" class="sr-only"> Link </Label>
                      <pre class="rounded-sm overflow-hidden pl-2 pr-3 py-3 bg-gray-100"><code>tailscale up --login-server https://wrt.12bit.xyz</code></pre>
                    </div>
                    <Button type="submit" size="sm" class="px-3">
                      <span class="sr-only">Copy</span>
                      <CopyIcon class="w-4 h-4" />
                    </Button>
                  </div>
                  <DialogFooter class="sm:justify-start"> </DialogFooter>
                </DialogContent>
              </Dialog>
            </div>
          </CardHeader>
          <CardContent>
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead class="md:w-1/3 flex-auto md:flex-initial md:shrink-0 w-0 text-ellipsis"> NODE </TableHead>
                  <TableHead class=""> ADDRESSES </TableHead>
                  <TableHead class=""> VERSION </TableHead>
                  <TableHead class="pl-2"> STATUS </TableHead>
                  <TableHead class=""> <span class="sr-only">Nodes action menu</span> </TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                <TableRow v-for="node in nodes" :key="node.id">
                  <TableCell class="w-[100px]">{{ node.name }}</TableCell>
                  <TableCell> {{ node.addr }} </TableCell>
                  <TableCell class=""> 1.78.1 </TableCell>
                  <TableCell>
                    <span class="inline-block w-2 h-2 rounded-full mr-1" :class="node.status ? 'bg-green-300 dark:bg-green-400' : 'bg-gray-300 dark:bg-gray-500'"></span>
                    {{ node.status ? 'online' : 'offline' }}
                  </TableCell>
                  <TableCell class="cursor-pointer">
                    <DropdownMenu>
                      <DropdownMenuTrigger as-child>
                        <Ellipsis class="h-5 w-5" />
                      </DropdownMenuTrigger>
                      <DropdownMenuContent align="end">
                        <DropdownMenuItem>Edit node name</DropdownMenuItem>
                        <DropdownMenuItem>Edit node IPv4</DropdownMenuItem>
                        <DropdownMenuItem>Disable key expiry</DropdownMenuItem>
                        <DropdownMenuSeparator />
                        <DropdownMenuItem>Edit ACL tags</DropdownMenuItem>
                        <DropdownMenuSeparator />
                        <DropdownMenuItem class="text-red-400">Remove...</DropdownMenuItem>
                      </DropdownMenuContent>
                    </DropdownMenu>
                  </TableCell>
                </TableRow>
              </TableBody>
            </Table>
          </CardContent>
        </Card>
      </div>
    </main>
  </div>
</template>
