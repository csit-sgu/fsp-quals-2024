<script setup lang="ts">
import { ScrollArea, ScrollBar } from '@/components/ui/scroll-area'
import { WeeklyEvent } from '@/components/ui'
import { Separator } from '@/components/ui/separator'

const weekdays = ['Понедельник', 'Вторник', 'Среда', 'Четверг', 'Пятница', 'Суббота', 'Воскресенье']
</script>

<template>
  <ScrollArea class="border-y">
    <div class="w-max max-w-[1500px] min-h-full grid grid-cols-7 gap-2 pt-2">
      <div v-for="weekday in weekdays" class="min-w-8 justify-self-center font-bold text-lg">
        {{ weekday }}
      </div>
      <Separator class="col-span-7" />

      <WeeklyEvent class="mx-2 col-start-2 col-span-4" />
      <div class="col-span-1"></div>

      <WeeklyEvent class="mx-2 col-span-7" />

      <WeeklyEvent class="mx-2 col-start-3 col-span-2" />
      <div class="col-span-2"></div>

      <WeeklyEvent class="mx-2 col-start-3 col-span-1" />
      <div class="col-span-3"></div>

      <WeeklyEvent class="mx-2 col-start-2 col-span-4" />
      <div class="col-span-1"></div>
    </div>
    <ScrollBar orientation="horizontal" />
  </ScrollArea>
</template>
