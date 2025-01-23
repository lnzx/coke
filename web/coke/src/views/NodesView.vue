<script setup lang="ts">
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu'
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog'
import { Plus, Ellipsis, CopyIcon } from 'lucide-vue-next'
import { ref, onMounted } from 'vue'

import AppLayout from '@/components/AppLayout.vue'

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
  <AppLayout>
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
                    <DialogDescription> 要将新节点添加到您的 Tailnet，只需在该设备上安装 Tailscale 并运行以下命令。 </DialogDescription>
                  </DialogHeader>
                  <div class="flex items-center space-x-2">
                    <div class="grid flex-1 gap-2">
                      <Label for="link" class="sr-only"> Link </Label>
                      <pre class="rounded-sm overflow-hidden pl-2 pr-3 py-3 bg-gray-100"><code>tailscale up --login-server https://x.12bit.xyz</code></pre>
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
  </AppLayout>
</template>
