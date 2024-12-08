<script setup lang="ts">
import { ref, type Ref } from 'vue'
import { CalendarDate } from '@internationalized/date'
import { RangeCalendar } from '@/components/ui/range-calendar'
import type { DateRange } from 'radix-vue'
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

const options = new Map([
  ['Ближайшая неделя', ['/weekly', 'week']],
  ['Ближайший месяц', ['/table', 'month']],
  ['Ближайший квартал', ['/table', 'quarter']],
  ['Ближайшие полгода', ['/table', 'half']],
  ['Указать вручную', ['/table', 'custom']],
])

const emit = defineEmits<{
  (e: 'update', path: string, value: string): void
  (e: 'calendarUpdate', range: DateRange): void
}>()

const value = ref('')
const open = ref(false)

const today = new Date()
const startDate = new CalendarDate(today.getFullYear(), today.getMonth() + 1, today.getDate())
const calendarPickedRange = ref({
  start: startDate,
  end: startDate.add({ days: 20 }),
}) as Ref<DateRange>

const calendarUpdated = (range: DateRange) => emit('calendarUpdate', range)

const MAX_LENGTH = 21

</script>

<template>
  <Popover v-model:open="open">
    <PopoverTrigger as-child>
      <Button variant="outline" role="combobox" :aria-expanded="open" class="w-full justify-between">
        {{ value ? (value.length > MAX_LENGTH ? value.slice(0, MAX_LENGTH - 2) + '...' : value)
          : "Выберите диапазон..." }}
        <ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
      </Button>
    </PopoverTrigger>
    <PopoverContent class="w-[19em] p-0">
      <Command>
        <CommandList>
          <CommandGroup>
            <CommandItem v-for="option in options" :value="option" @select="(ev) => {
              if (typeof ev.detail.value === 'object') {
                value = ev.detail.value[0]
                emit('update', ev.detail.value[1][0], ev.detail.value[1][1])
              }
              console.log(typeof ev.detail.value)
              open = false
            }">
              {{ option[0] }}
              <Check :class="cn(
                'ml-auto h-4 w-4',
                value === option[0] ? 'opacity-100' : 'opacity-0',
              )" />
            </CommandItem>
          </CommandGroup>
        </CommandList>
      </Command>
    </PopoverContent>
  </Popover>
  <RangeCalendar v-if="value == 'Указать вручную'" v-model="calendarPickedRange" initial-focus
    @update:modelValue="calendarUpdated" />
</template>
