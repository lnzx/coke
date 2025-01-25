<script setup>
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { Switch } from '@/components/ui/switch'
import { Card, CardContent, CardDescription, CardHeader, CardTitle, CardFooter } from '@/components/ui/card'
import SettingsLayout from '@/components/settings/SettingsLayout.vue'
import { ref } from 'vue'

const disableMagicDns = ref(true)
</script>

<template>
  <SettingsLayout>
    <Card>
      <CardHeader>
        <CardTitle>DNS</CardTitle>
        <CardDescription> Tailnet name </CardDescription>
      </CardHeader>
      <CardContent>
        <form class="flex flex-col gap-4">
          <Input placeholder="dns" />
        </form>
      </CardContent>
      <CardFooter class="border-t px-6 py-4">
        <Button>Save</Button>
      </CardFooter>
    </Card>
    <Card>
      <CardHeader>
        <CardTitle>Nameservers</CardTitle>
        <CardDescription> Set the nameservers used by devices on your network to resolve DNS queries. </CardDescription>
      </CardHeader>
      <CardContent>
        <form class="flex flex-col gap-4">
          <Input placeholder="nameservers" default-value="100.100.100.100" />
        </form>
      </CardContent>
      <CardFooter class="border-t px-6 py-4">
        <Button>Save</Button>
      </CardFooter>
    </Card>
    <Card>
      <CardHeader>
        <CardTitle>MagicDNS</CardTitle>
        <CardDescription>Automatically register domain names for devices in your tailnet. This lets you to use a machine’s name instead of its IP address.</CardDescription>
      </CardHeader>
      <CardContent>
        <div class="flex items-center space-x-2">
          <Switch id="dis-magic" :checked="disableMagicDns" />
          <label for="include" class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"> Disable MagicDNS </label>
        </div>
      </CardContent>
    </Card>
  </SettingsLayout>
</template>
