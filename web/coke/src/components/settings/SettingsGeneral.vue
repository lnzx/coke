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
        <CardTitle>Domain</CardTitle>
        <CardDescription> The url clients will connect to. </CardDescription>
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
  </SettingsLayout>
</template>
