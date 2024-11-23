<script setup lang="ts">
import { Button } from '@/components/ui/button'
import {
  Command,
  CommandGroup,
  CommandItem,
  CommandEmpty,
  CommandInput,
  CommandList,
} from '@/components/ui/command'
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@/components/ui/popover'
import { cn } from '@/lib/utils'
import { Check, ChevronsUpDown } from 'lucide-vue-next'
import { ref } from 'vue'

type Props = {
  options: string[],
  defaultMsg: string,
  showSearch?: boolean,
}

const props = defineProps<Props>()

const emit = defineEmits<{
  (e: 'update', value: string): void
}>()

const MAX_LENGTH = 21

const value = ref('')
const open = ref(false)
</script>

<template>
  <Popover v-model:open="open">
    <PopoverTrigger as-child>
      <Button variant="outline" role="combobox" :aria-expanded="open" class="w-[250px] justify-between">
        {{ value ? (value.length > MAX_LENGTH ? value.slice(0, MAX_LENGTH-2) + '...' : value)
            : defaultMsg }}
        <ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
      </Button>
    </PopoverTrigger>
    <PopoverContent class="w-[250px] p-0">
      <Command>
        <CommandInput v-if="showSearch" class="h-9" placeholder="Поиск..." />
        <CommandEmpty v-if="showSearch">Ничего не найдено.</CommandEmpty>
        <CommandList>
          <CommandGroup>
            <CommandItem v-for="option in props.options.filter(x => x.length > 0)" :key="option" :value="option" @select="(ev) => {
              if (typeof ev.detail.value === 'string') {
                value = ev.detail.value
                emit('update', value)
              }
              open = false
            }">
              {{ option }}
              <Check :class="cn(
                'ml-auto h-4 w-4',
                value === option ? 'opacity-100' : 'opacity-0',
              )" />
            </CommandItem>
          </CommandGroup>
        </CommandList>
      </Command>
    </PopoverContent>
  </Popover>
</template>
