<script setup lang="ts">
import { Button } from '@/components/ui/button'
import { CardFooter } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'

import AppLayout from '@/components/AppLayout.vue'
</script>

<template>
  <AppLayout>
    <main class="flex min-h-[calc(100vh_-_theme(spacing.16))] flex-1 flex-col gap-4 p-4 md:gap-8 md:p-10">
      <div class="mx-auto grid w-full max-w-6xl gap-2">
        <h1 class="text-3xl font-semibold">Settings</h1>
      </div>
      <div class="mx-auto grid w-full max-w-6xl items-start gap-6 md:grid-cols-[180px_1fr] lg:grid-cols-[250px_1fr]">
        <nav class="grid gap-4 text-sm text-muted-foreground">
          <a href="#" class="font-semibold text-primary"> General </a>
          <a href="#"> Keys </a>
          <a href="#"> Node management </a>
          <a href="#"> DNS </a>
        </nav>
        <div class="grid gap-6">
          <Card>
            <CardHeader>
              <CardTitle>Domain</CardTitle>
              <CardDescription> Authenticate nodes without an interactive login. </CardDescription>
            </CardHeader>
            <CardContent>
              <form>
                <Input placeholder="https://example.com:443" />
              </form>
            </CardContent>
            <CardFooter class="border-t px-6 py-4">
              <Button>Save</Button>
            </CardFooter>
          </Card>
          <Card>
            <CardHeader>
              <CardTitle>Auth keys</CardTitle>
              <CardDescription> Authenticate nodes without an interactive login. </CardDescription>
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
              <CardDescription> Access tokens give access to the Headscale API. </CardDescription>
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
        </div>
      </div>
    </main>
  </AppLayout>
</template>
