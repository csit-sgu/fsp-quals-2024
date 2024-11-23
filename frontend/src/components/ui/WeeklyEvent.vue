<script setup lang="ts">
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card'
import { cn } from '@/lib/utils'
import type { HTMLAttributes } from 'vue'
import { type PrimitiveProps } from 'radix-vue'
import type { Competition } from '@/lib/dataSource'

interface Props extends PrimitiveProps {
  eventData: Competition,
  class?: HTMLAttributes['class']
}

const props = defineProps<Props>()
const eventData = props.eventData
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
          <div v-for="{ country, locality } in eventData.location_data">
            <span>{{ country }}, {{ locality }}</span>
          </div>
          <span>{{ eventData.n_participants }} человек</span>
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
          <div v-for="{ country, region, locality } in eventData.location_data">
            <span>{{ country }}</span>
            <span v-if="region !== null">, {{ region }}</span>
            <span>, {{ locality }}</span>
          </div>
        </div>
        <div class="grid grid-cols-2 gap-0">
          <div class="font-semibold">Начало соревнования:</div>
          <div>{{ eventData.start_date }}</div>

          <div class="font-semibold">Конец соревнования:</div>
          <div>{{ eventData.end_date }}</div>

          <div class="font-semibold">Количество участников:</div>
          <div>{{ eventData.n_participants }}</div>
        </div>

        <div class="flex flex-col">
          <div class="font-semibold">Допустимые пол и возрастная группа:</div>
          <div v-for="{ gender, left_bound, right_bound, original } in eventData.age_data">
            <div v-if="left_bound != right_bound">
              {{ original || gender }}
              <span v-if="left_bound"> от {{ left_bound }} лет</span>
              <span v-if="right_bound < 255"> до {{ right_bound }} лет</span>
            </div>
            <div v-if="left_bound === right_bound">
              {{ original || gender }}, {{ left_bound }} лет
            </div>
          </div>
        </div>

        <div>{{ eventData.additional_info }}</div>
      </DialogHeader>
    </DialogContent>
  </Dialog>
</template>
