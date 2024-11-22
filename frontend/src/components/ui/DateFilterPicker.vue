<script setup lang="ts">
import { Chooser } from '@/components/ui'
import { ref, type Ref } from 'vue'
import { CalendarDate } from '@internationalized/date'
import { RangeCalendar } from '@/components/ui/range-calendar'
import type { DateRange } from 'radix-vue'

const options = [
  { value: 'week', label: 'Ближайшая неделя' },
  { value: 'month', label: 'Ближайший месяц' },
  { value: 'quarter', label: 'Ближайший квартал' },
  { value: 'half-year', label: 'Ближайшие полгода' },
  { value: 'manual', label: 'Указать вручную' },
]

const value = ref('')

const today = new Date();
const startDate = new CalendarDate(today.getFullYear(), today.getMonth(), today.getDay());
const calendarPickedRange = ref({
  start: startDate,
  end: startDate.add({ days: 20 }),
}) as Ref<DateRange>;

const calendarUpdated = (range: DateRange) => {
  if (!range.start || !range.end) {
    return;
  }
  console.log(range.start, range.end);
};

const update = (newValue: string) => {
  value.value = newValue
};
</script>

<template>
  <Chooser :options="options" default-msg="Выберите диапазон..." @update="update" />
  <RangeCalendar v-if="value == 'manual'" v-model="calendarPickedRange" initial-focus
    @update:modelValue="calendarUpdated" />
</template>
