<script setup lang="ts">
import { ScrollArea, ScrollBar } from '@/components/ui/scroll-area'
</script>

<template>
  <ScrollArea class="border-2 border-black w-full">
    <div class="w-max grid grid-cols-7 overflow-clip h-full gap-2 border-red-500 border-2">
      <div class="border-black border-2 col-span-7">some competition</div>

      <div class="border-black border-2 col-start-3 col-span-2">some competition</div>
      <div class="col-span-2"></div>

      <div class="border-black border-2 col-start-2 col-span-4">some competition</div>
    </div>
    <ScrollBar orientation="horizontal" />
  </ScrollArea>
</template>
