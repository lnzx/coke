<script setup>
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { ref } from 'vue'
import { useApi } from '@/composable/useApi'

const api = useApi()
const email = ref('')
const password = ref('')

const login = () => {
  if (!email.value || !password.value) {
    return
  }
  api
    .post('/api/login', {
      username: email.value,
      password: password.value,
    })
    .then((rsp) => {
      const token = rsp.data.token
      if (token) {
        localStorage.setItem('token', token)
        location.href = '/nodes'
      }
    })
}
</script>

<template>
  <Card class="w-full max-w-sm">
    <CardHeader>
      <CardTitle class="text-2xl"> Login to Headscale </CardTitle>
      <CardDescription> Enter your email below to login to your account. </CardDescription>
    </CardHeader>
    <CardContent class="grid gap-4">
      <div class="grid gap-2">
        <Label for="email">Email</Label>
        <Input id="email" type="email" placeholder="m@example.com" v-model="email" required />
      </div>
      <div class="grid gap-2">
        <Label for="password">Password</Label>
        <Input id="password" type="password" v-model="password" required />
      </div>
    </CardContent>
    <CardFooter>
      <Button class="w-full" @click="login"> Sign in </Button>
    </CardFooter>
  </Card>
</template>
