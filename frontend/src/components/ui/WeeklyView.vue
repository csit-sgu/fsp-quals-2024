<script setup lang="ts">
import { ScrollArea, ScrollBar } from '@/components/ui/scroll-area'
import { WeeklyEvent } from '@/components/ui'
import { Separator } from '@/components/ui/separator'
import { cn } from '@/lib/utils'

const weekdays = ['Понедельник', 'Вторник', 'Среда', 'Четверг', 'Пятница', 'Суббота', 'Воскресенье']

const events = [
  { duration: 4, startDay: 'Вторник' },
  { duration: 7, startDay: 'Понедельник' },
  { duration: 2, startDay: 'Среда' },
  { duration: 1, startDay: 'Среда' },
  { duration: 4, startDay: 'Вторник' },
]

const renderEvents = events.map(({ duration, startDay }) => {
  const startDayIdx = weekdays.indexOf(startDay)
  const clampedDuration = Math.min(duration, weekdays.length - startDayIdx)
  return {
    classes: [
      'col-start-' + (startDayIdx + 1).toString(),
      'col-span-' + clampedDuration.toString(),
    ],
    tailDays: weekdays.length - clampedDuration - startDayIdx,
  }
})
</script>

<template>
  <ScrollArea class="border-t pb-4">
    <div class="w-max max-w-[1500px] min-h-full grid grid-cols-7 gap-2 pt-2">
      <div v-for="weekday in weekdays" class="min-w-8 justify-self-center font-bold text-lg">
        {{ weekday }}
      </div>
      <Separator class="col-span-7" />

      <template v-for="{ classes, tailDays } in renderEvents">
        <WeeklyEvent :class="cn('mx-2', ...classes)" />
        <div v-if="tailDays > 0" :class="'col-span-' + tailDays.toString()"></div>
      </template>
    </div>
    <ScrollBar orientation="horizontal" />
  </ScrollArea>
</template>
