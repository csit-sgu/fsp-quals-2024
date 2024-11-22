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
    classes: ['colspan' + clampedDuration.toString()],
    startDays: startDayIdx,
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
      <Separator class="colspan7" />

      <template v-for="{ classes, startDays, tailDays } in renderEvents">
        <div v-if="startDays > 0" :class="'colspan' + startDays.toString()"></div>
        <WeeklyEvent :class="cn('mx-2', ...classes)" />
        <div v-if="tailDays > 0" :class="'colspan' + tailDays.toString()"></div>
      </template>
    </div>
    <ScrollBar orientation="horizontal" />
  </ScrollArea>
</template>

<style>
/* NOTE(aguschin): for whatewer reason tailwind's col-span-* do not work, but this works */
.colspan7 {
  grid-column: span 7 / span 7;
}
.colspan6 {
  grid-column: span 6 / span 6;
}
.colspan5 {
  grid-column: span 5 / span 5;
}
.colspan4 {
  grid-column: span 4 / span 4;
}
.colspan3 {
  grid-column: span 3 / span 3;
}
.colspan2 {
  grid-column: span 2 / span 2;
}
.colspan1 {
  grid-column: span 1 / span 1;
}
</style>
