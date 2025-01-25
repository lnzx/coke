<script setup>
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuSeparator, DropdownMenuTrigger } from '@/components/ui/dropdown-menu'
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table'
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog'
import { Plus, Ellipsis, CopyIcon } from 'lucide-vue-next'
import AppLayout from '@/components/AppLayout.vue'
import { ref } from 'vue'
import { useApi } from '@/composable/useApi'
import dayjs from 'dayjs'

const api = useApi()
const users = ref([])
const getUsers = () => {
  api.get('/api/users').then((rsp) => {
    users.value = rsp.data
  })
}

const formatTimestamp = (seconds) => {
  return dayjs.unix(seconds).format('YYYY-MM-DD HH:mm:ss') // 自定义格式
}

getUsers()
</script>

<template>
  <AppLayout>
    <div class="grid gap-4 md:gap-8">
      <Card class="xl:col-span-2">
        <CardHeader class="flex flex-row items-center">
          <div class="grid gap-2">
            <CardTitle class="text-2xl font-semibold">Users</CardTitle>
            <CardDescription>
              <Badge variant="secondary">{{ users.length }} users</Badge>
            </CardDescription>
          </div>
          <div class="ml-auto gap-1">
            <Dialog class="ml-auto gap-1">
              <DialogTrigger as-child>
                <Button size="sm"> Add user <Plus class="h-4 w-4" /> </Button>
              </DialogTrigger>
              <DialogContent class="sm:max-w-xl">
                <DialogHeader>
                  <DialogTitle>Add user</DialogTitle>
                  <DialogDescription> To add a new device to your tailnet, simply install Tailscale on that device and run the following command</DialogDescription>
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
                <TableHead class="w-1/8"> ID </TableHead>
                <TableHead class="w-1/4"> NAME </TableHead>
                <TableHead class="w-1/4"> CREATED_AT </TableHead>
                <TableHead class="w-1/4 text-right"> <span class="sr-only">Nodes action menu</span> </TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              <TableRow v-for="user in users" :key="user.id">
                <TableCell class="w-1/8"> {{ user.id }} </TableCell>
                <TableCell class="w-1/4">{{ user.name }}</TableCell>
                <TableCell class="w-1/4"> {{ formatTimestamp(user.created_at.seconds) }} </TableCell>
                <TableCell class="w-1/4 text-right cursor-pointer">
                  <DropdownMenu>
                    <DropdownMenuTrigger as-child>
                      <Ellipsis class="h-5 w-5" />
                    </DropdownMenuTrigger>
                    <DropdownMenuContent align="end">
                      <DropdownMenuItem>Edit user name</DropdownMenuItem>
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
  </AppLayout>
</template>
