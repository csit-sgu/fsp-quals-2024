<script setup lang="ts">
import { ScrollArea, ScrollBar } from '@/components/ui/scroll-area'
import { WeeklyEvent } from '@/components/ui'
import { Separator } from '@/components/ui/separator'
import { cn } from '@/lib/utils'
import dayjs from 'dayjs'
import isoWeek from 'dayjs/plugin/isoWeek'
import utc from 'dayjs/plugin/utc'

const props = defineProps<{
  beginDay?: string
}>()

dayjs.extend(isoWeek)
dayjs.extend(utc)

const weekdays = ['Понедельник', 'Вторник', 'Среда', 'Четверг', 'Пятница', 'Суббота', 'Воскресенье']

const events = [
  {
    code: '2191860017026583',
    start_date: '2024-11-24T00:00:00Z',
    location_data: [
      {
        country: 'РОССИЯ',
        region: '',
        locality: 'поселок городского типа',
      },
    ],
    age_data: [
      {
        gender: 'male',
        left_bound: 0,
        right_bound: 0,
        original: '',
      },
    ],
    title: 'ЧЕМПИОНАТ ФЕДЕРАЛЬНОГО ОКРУГА (УРАЛЬСКИЙ ФЕДЕРАЛЬНЫЙ ОКРУГ)',
    additional_info: 'РИТМ - СИМУЛЯТОР, ДВОЕБОРЬЕ - ТАКТИЧЕСКАЯ СТРЕЛЬБА',
    n_participants: 320,
    stage: '',
    end_date: '2024-11-26T00:00:00Z',
    sport: '',
  },
  {
    code: '2191860017026583',
    start_date: '2024-11-23T00:00:00Z',
    location_data: [
      {
        country: 'РОССИЯ',
        region: '',
        locality: 'поселок городского типа',
      },
    ],
    age_data: [
      {
        gender: 'male',
        left_bound: 0,
        right_bound: 0,
        original: '',
      },
    ],
    title: 'ЧЕМПИОНАТ ФЕДЕРАЛЬНОГО ОКРУГА (УРАЛЬСКИЙ ФЕДЕРАЛЬНЫЙ ОКРУГ)',
    additional_info: 'РИТМ - СИМУЛЯТОР, ДВОЕБОРЬЕ - ТАКТИЧЕСКАЯ СТРЕЛЬБА',
    n_participants: 320,
    stage: '',
    end_date: '2024-11-27T00:00:00Z',
    sport: '',
  },
  {
    code: '2191860017026583',
    start_date: '2024-11-25T00:00:00Z',
    location_data: [
      {
        country: 'РОССИЯ',
        region: '',
        locality: 'поселок городского типа',
      },
    ],
    age_data: [
      {
        gender: 'male',
        left_bound: 0,
        right_bound: 0,
        original: '',
      },
    ],
    title: 'ЧЕМПИОНАТ ФЕДЕРАЛЬНОГО ОКРУГА (УРАЛЬСКИЙ ФЕДЕРАЛЬНЫЙ ОКРУГ)',
    additional_info: 'РИТМ - СИМУЛЯТОР, ДВОЕБОРЬЕ - ТАКТИЧЕСКАЯ СТРЕЛЬБА',
    n_participants: 320,
    stage: '',
    end_date: '2024-11-27T00:00:00Z',
    sport: '',
  },
  {
    code: '2191860017026583',
    start_date: '2024-11-25T00:00:00Z',
    location_data: [
      {
        country: 'РОССИЯ',
        region: '',
        locality: 'поселок городского типа',
      },
    ],
    age_data: [
      {
        gender: 'male',
        left_bound: 0,
        right_bound: 0,
        original: '',
      },
    ],
    title: 'ЧЕМПИОНАТ ФЕДЕРАЛЬНОГО ОКРУГА (УРАЛЬСКИЙ ФЕДЕРАЛЬНЫЙ ОКРУГ)',
    additional_info: 'РИТМ - СИМУЛЯТОР, ДВОЕБОРЬЕ - ТАКТИЧЕСКАЯ СТРЕЛЬБА',
    n_participants: 320,
    stage: '',
    end_date: '2024-12-03T00:00:00Z',
    sport: '',
  },
  {
    code: '2191860017026583',
    start_date: '2024-11-20T00:00:00Z',
    location_data: [
      {
        country: 'РОССИЯ',
        region: '',
        locality: 'поселок городского типа',
      },
    ],
    age_data: [
      {
        gender: 'male',
        left_bound: 0,
        right_bound: 0,
        original: '',
      },
    ],
    title: 'ЧЕМПИОНАТ ФЕДЕРАЛЬНОГО ОКРУГА (УРАЛЬСКИЙ ФЕДЕРАЛЬНЫЙ ОКРУГ)',
    additional_info: 'РИТМ - СИМУЛЯТОР, ДВОЕБОРЬЕ - ТАКТИЧЕСКАЯ СТРЕЛЬБА',
    n_participants: 320,
    stage: '',
    end_date: '2024-12-03T00:00:00Z',
    sport: '',
  },
  {
    code: '2191860017026583',
    start_date: '2024-11-20T00:00:00Z',
    location_data: [
      {
        country: 'РОССИЯ',
        region: '',
        locality: 'поселок городского типа',
      },
    ],
    age_data: [
      {
        gender: 'male',
        left_bound: 0,
        right_bound: 0,
        original: '',
      },
    ],
    title: 'ЧЕМПИОНАТ ФЕДЕРАЛЬНОГО ОКРУГА (УРАЛЬСКИЙ ФЕДЕРАЛЬНЫЙ ОКРУГ)',
    additional_info: 'РИТМ - СИМУЛЯТОР, ДВОЕБОРЬЕ - ТАКТИЧЕСКАЯ СТРЕЛЬБА',
    n_participants: 320,
    stage: '',
    end_date: '2024-11-24T00:00:00Z',
    sport: '',
  },
]

const today = dayjs.utc(props.beginDay).set('hour', 13) || dayjs.utc().set('hour', 13)
const beginWeek = today.isoWeekday() - 1

const renderEvents = events
  .map(({ start_date, end_date }) => {
    const start = dayjs(start_date)
    const end = dayjs(end_date)
    const startDiffToday = Math.round(start.diff(today, 'day', true)) + 1
    if (startDiffToday > weekdays.length) {
      return null
    }

    const endDiffToday = Math.round(end.diff(today, 'day', true)) + 1

    const headDays = Math.max(0, startDiffToday)
    const renderDuration =
      endDiffToday > weekdays.length ? weekdays.length - headDays : endDiffToday - headDays + 1
    const tailDays = weekdays.length - renderDuration - headDays

    return {
      classes: ['colspan' + renderDuration.toString()],
      headDays: headDays,
      tailDays: tailDays,
    }
  })
  .filter((res) => res !== null)

const adjustedWeek = weekdays.slice(beginWeek).concat(weekdays.slice(0, beginWeek))
</script>

<template>
  <ScrollArea class="border-t pb-4">
    <div class="w-max max-w-[1500px] min-h-full grid grid-cols-7 gap-2 pt-2">
      <div v-for="weekday in adjustedWeek" class="min-w-8 justify-self-center font-bold text-lg">
        {{ weekday }}
      </div>
      <Separator class="colspan7" />

      <template v-for="{ classes, headDays, tailDays } in renderEvents">
        <div v-if="headDays > 0" :class="'colspan' + headDays.toString()"></div>
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
