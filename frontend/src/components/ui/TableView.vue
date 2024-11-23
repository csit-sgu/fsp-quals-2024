<script setup lang="ts">
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { Separator } from '@/components/ui/separator'
import { type Competition } from '@/lib/dataSource'
import { onMounted, ref, type Ref } from 'vue'

const props = defineProps<{
  eventsPromise: Promise<Competition[]>,
}>()

const events: Ref<Competition[]> = ref([])
onMounted(async () => {
  events.value = await props.eventsPromise
})
</script>

<template>
  <Table>
    <TableHeader>
      <TableRow>
        <TableHead class="text-center w-48">Вид спорта</TableHead>
        <TableHead class="text-center w-96">Название и дисциплина/программа</TableHead>
        <TableHead class="text-center w-48">Этап</TableHead>
        <TableHead class="text-center w-48">Пол и возрастные категории</TableHead>
        <TableHead class="text-center w-48">Сроки проведения</TableHead>
        <TableHead class="text-center w-48">Место проведения</TableHead>
        <TableHead class="text-center w-20">Количество участников</TableHead>
      </TableRow>
    </TableHeader>
    <TableBody>
      <TableRow v-for="event in events">
        <!-- TODO(aguschin): remove fallback string -->
        <TableCell class="text-center">{{ event.sport || 'Не указано' }}</TableCell>
        <TableCell class="text-left flex flex-col">
          <div class="font-semibold">{{ event.title }}</div>
          <div class="font-light">{{ event.additional_info }}</div>
        </TableCell>
        <!-- TODO(aguschin): remove fallback string -->
        <TableCell class="text-center">{{ event.stage || 'Не указано' }}</TableCell>
        <TableCell class="text-center">
          <div class="flex flex-col">
            <div class="font-semibold"></div>
            <div v-for="{ gender, left_bound, right_bound, original } in event.age_data">
              <!-- TODO(aguschin): use only original string? -->
              <span>{{ original || gender }}</span>
              <span v-if="left_bound !== 0"> от {{ left_bound }} лет</span>
              <span v-if="right_bound !== 0"> до {{ right_bound }} лет</span>
            </div>
          </div>
        </TableCell>
        <TableCell>
          <div class="flex flex-col gap-0">
            <div class="font-semibold">Начало:</div>
            <div>
              {{
                new Date(event.start_date).toLocaleDateString('ru-RU', {
                  day: 'numeric',
                  month: 'long',
                  year: 'numeric',
                })
              }}
            </div>

            <div class="font-semibold">Окончание:</div>
            <div>
              {{
                new Date(event.end_date).toLocaleDateString('ru-RU', {
                  day: 'numeric',
                  month: 'long',
                  year: 'numeric',
                })
              }}
            </div>
          </div>
        </TableCell>
        <TableCell>
          <div v-for="{ country, region, locality } in event.location_data" class="flex flex-col">
            <span>{{ country }},</span>
            <span v-if="region">{{ region }},</span>
            <span>{{ locality }}</span>
            <Separator v-if="event.location_data.length > 1" />
          </div>
        </TableCell>
        <TableCell class="text-center">{{ event.n_participants }}</TableCell>
      </TableRow>
    </TableBody>
  </Table>
</template>
