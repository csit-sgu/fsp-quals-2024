<script setup lang="ts">
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { Button } from './button'
</script>

<template>
  <Card>
    <CardHeader class="px-4 pt-3 pb-1">
      <CardTitle class="text-xl">ЧЕМПИОНАТ ЦЕНТРАЛЬНОГО ФЕДЕРАЛЬНОГО ОКРУГА</CardTitle>
      <CardDescription>АВИАМОДЕЛЬНЫЙ СПОРТ</CardDescription>
    </CardHeader>
    <CardContent class="px-4 pt-0 pb-3 flex flex-col gap-x-5 text-sm">
      <!-- <span>АВИАМОДЕЛЬНЫЙ СПОРТ</span> -->
      <span>РОССИЯ, г. Орехово-Зуево</span>
      <span>25 человек</span>
    </CardContent>
    <!-- <CardFooter class="flex gap-x-5 px-6 pb-6">
      <Button>Открыть подробности</Button>
    </CardFooter> -->
  </Card>
</template>
