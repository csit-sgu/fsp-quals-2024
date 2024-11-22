<script setup lang="ts">
import { Button } from '@/components/ui/button'
import {
  Command,
  CommandGroup,
  CommandItem,
  CommandList,
} from '@/components/ui/command'

import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover'
import { cn } from '@/lib/utils'
import { Check, ChevronsUpDown } from 'lucide-vue-next'
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

const value = ref('');
const open = ref(false)

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
</script>

<template>
  <Popover v-model:open="open">
    <PopoverTrigger as-child>
      <Button variant="outline" role="combobox" :aria-expanded="open" class="w-[200px] justify-between">
        {{ value
          ? options.find((x) => x.value === value)?.label
          : "Выберите диапазон..." }}
        <ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
      </Button>
    </PopoverTrigger>
    <PopoverContent class="w-[200px] p-0">
      <Command>
        <CommandList>
          <CommandGroup>
            <CommandItem v-for="option in options" :key="option.value" :value="option.value" @select="(ev) => {
              if (typeof ev.detail.value === 'string') {
                value = ev.detail.value
              }
              open = false
            }">
              {{ option.label }}
              <Check :class="cn(
                'ml-auto h-4 w-4',
                value === option.value ? 'opacity-100' : 'opacity-0',
              )" />
            </CommandItem>
          </CommandGroup>
        </CommandList>
      </Command>
    </PopoverContent>
  </Popover>

  <RangeCalendar v-if="value == 'manual'" v-model="calendarPickedRange" initial-focus
    @update:modelValue="calendarUpdated" />
</template>
