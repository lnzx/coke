<script setup>
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle, CardFooter } from '@/components/ui/card'
import SettingsLayout from '@/components/settings/SettingsLayout.vue'
</script>

<template>
  <SettingsLayout>
    <Card>
      <CardHeader>
        <CardTitle>Auth keys</CardTitle>
        <CardDescription> Authenticate devices without an interactive login. </CardDescription>
      </CardHeader>
      <CardContent>
        <form>
          <Input placeholder="auth key" />
        </form>
      </CardContent>
      <CardFooter class="border-t px-6 py-4">
        <Button>Generate auth key…</Button>
      </CardFooter>
    </Card>
    <Card>
      <CardHeader>
        <CardTitle>API access tokens</CardTitle>
        <CardDescription> Access tokens give access to the Tailscale API. </CardDescription>
      </CardHeader>
      <CardContent>
        <form class="flex flex-col gap-4">
          <Input placeholder="api access token" />
        </form>
      </CardContent>
      <CardFooter class="border-t px-6 py-4">
        <Button>Generate access token…</Button>
      </CardFooter>
    </Card>
  </SettingsLayout>
</template>
