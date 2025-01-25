<script setup>
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Checkbox } from '@/components/ui/checkbox'
import { Switch } from '@/components/ui/switch'
import { Card, CardContent, CardDescription, CardHeader, CardTitle, CardFooter } from '@/components/ui/card'
import SettingsLayout from '@/components/settings/SettingsLayout.vue'
</script>

<template>
  <SettingsLayout>
    <Card>
      <CardHeader>
        <CardTitle>Device management</CardTitle>
        <CardDescription> Manage your tailnet device settings. </CardDescription>
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
    <Card>
      <CardHeader>
        <CardTitle>Auto-update Tailscale</CardTitle>
      </CardHeader>
      <CardContent>
        <div class="flex items-center space-x-2">
          <Switch id="airplane-mode" />
          <label for="include" class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"> Auto-update new devices </label>
        </div>
      </CardContent>
    </Card>
  </SettingsLayout>
</template>
