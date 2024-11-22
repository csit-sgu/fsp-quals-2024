<script setup lang="ts">
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { cn } from '@/lib/utils'
import type { HTMLAttributes } from 'vue'
import { Primitive, type PrimitiveProps } from 'radix-vue'
import { Button } from './button'

interface Props extends PrimitiveProps {
  class?: HTMLAttributes['class']
}

const props = defineProps<Props>()

type AgeGroup = {
  gender: string
  ages: string
}

const eventData = {
  code: '100003',
  sport: 'АВИАМОДЕЛЬНЫЙ СПОРТ',
  title: 'ЧЕМПИОНАТ ЦЕНТРАЛЬНОГО ФЕДЕРАЛЬНОГО ОКРУГА',
  locationData: [
    {
      country: 'РОССИЯ',
      region: 'МОСКОВСКАЯ ОБЛАСТЬ',
      locality: 'г. Орехово-Зуево',
    },
    {
      country: 'КАНАДА',
      region: 'АЛЬБЕРТА',
      locality: 'г. Калгари',
    },
  ],
  nParticipants: 25,
  startDate: '2024-02-16',
  endDate: '2024-02-19',
  ageGroups: [
    { gender: 'женщины', leftBound: 14, rightBound: null },
    { gender: 'мужчины', leftBound: null, rightBound: 20 },
    { gender: 'юниорки', leftBound: 14, rightBound: 20 },
  ],
  additionalInfo: 'С препятствиями, до 90 кг',
}
</script>

<template>
  <Dialog>
    <DialogTrigger :class="cn(props.class, 'text-start')">
      <Card>
        <CardHeader class="px-4 pt-3 pb-1">
          <CardTitle class="text-xl">{{ eventData.title }}</CardTitle>
          <CardDescription>{{ eventData.sport }}</CardDescription>
        </CardHeader>
        <CardContent class="px-4 pt-0 pb-3 flex flex-col gap-x-5 text-sm">
          <div v-for="{ country, locality } in eventData.locationData">
            <span>{{ country }}, {{ locality }}</span>
          </div>
          <span>{{ eventData.nParticipants }} человек</span>
        </CardContent>
      </Card>
    </DialogTrigger>
    <DialogContent>
      <DialogHeader class="px-4 pt-3 pb-1">
        <DialogTitle class="text-2xl">{{ eventData.title }}</DialogTitle>
        <DialogDescription>
          <div class="text-lg pb-0">{{ eventData.sport }}</div>
          <div class="text-sm">{{ eventData.code }}</div>
        </DialogDescription>
        <div class="py-0">
          <div v-for="{ country, region, locality } in eventData.locationData">
            <span>{{ country }}</span>
            <span v-if="region !== null">, {{ region }}</span>
            <span>, {{ locality }}</span>
          </div>
        </div>
        <div class="grid grid-cols-2 gap-0">
          <div class="font-semibold">Начало соревнования:</div>
          <div>{{ eventData.startDate }}</div>

          <div class="font-semibold">Конец соревнования:</div>
          <div>{{ eventData.endDate }}</div>

          <div class="font-semibold">Количество участников:</div>
          <div>{{ eventData.nParticipants }}</div>
        </div>

        <div class="flex flex-col">
          <div class="font-semibold">Допустимые пол и возрастная группа:</div>
          <div v-for="{ gender, leftBound, rightBound } in eventData.ageGroups">
            <span>{{ gender }}</span>
            <span v-if="leftBound !== null"> от {{ leftBound }} лет</span>
            <span v-if="rightBound !== null"> до {{ rightBound }} лет</span>
          </div>
        </div>

        <div>{{ eventData.additionalInfo }}</div>
      </DialogHeader>
    </DialogContent>
  </Dialog>
</template>
