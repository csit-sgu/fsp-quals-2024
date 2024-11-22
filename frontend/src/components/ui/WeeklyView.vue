<script setup lang="ts">
import { ScrollArea, ScrollBar } from '@/components/ui/scroll-area'
import { WeeklyEvent } from '@/components/ui'
import { Separator } from '@/components/ui/separator'
</script>

<template>
  <ScrollArea class="border-2 p-3 max-w-[1000px] w-full">
    <!-- <div
      class="w-max max-w-[1500px] grid grid-cols-7 overflow-clip h-full gap-2 border-red-500 border-2"
    > -->
    <div class="w-max max-w-[1500px] grid grid-cols-7 overflow-clip h-full gap-2">
      <div class="justify-self-center font-bold text-lg">Понедельник</div>
      <div class="justify-self-center font-bold text-lg">Вторник</div>
      <div class="justify-self-center font-bold text-lg">Среда</div>
      <div class="justify-self-center font-bold text-lg">Четверг</div>
      <div class="justify-self-center font-bold text-lg">Пятница</div>
      <div class="justify-self-center font-bold text-lg">Суббота</div>
      <div class="justify-self-center font-bold text-lg">Воскресенье</div>
      <Separator class="col-span-7" />

      <WeeklyEvent class="col-start-2 col-span-4" />
      <div class="col-span-1"></div>

      <WeeklyEvent class="col-span-7" />

      <WeeklyEvent class="col-start-3 col-span-2" />
      <div class="col-span-2"></div>

      <WeeklyEvent class="col-start-3 col-span-1" />
      <div class="col-span-3"></div>

      <WeeklyEvent class="col-start-2 col-span-4" />
      <div class="col-span-1"></div>
    </div>
    <ScrollBar orientation="horizontal" />
  </ScrollArea>
</template>
