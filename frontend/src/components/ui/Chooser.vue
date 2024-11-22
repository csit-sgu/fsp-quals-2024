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
import { ref } from 'vue'

type Option = { value: string, label: string }

type Props = {
  options: Option[],
  defaultMsg: string,
}

const props = defineProps<Props>()

const emit = defineEmits<{
  (e: 'update', value: string): void
}>()

const value = ref('')
const open = ref(false)
</script>

<template>
  <Popover v-model:open="open">
    <PopoverTrigger as-child>
      <Button variant="outline" role="combobox" :aria-expanded="open" class="w-[200px] justify-between">
        {{ value ? props.options.find((x) => x.value === value)?.label : defaultMsg }}
        <ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
      </Button>
    </PopoverTrigger>
    <PopoverContent class="w-[200px] p-0">
      <Command>
        <CommandList>
          <CommandGroup>
            <CommandItem v-for="option in props.options" :key="option.value" :value="option.value" @select="(ev) => {
              if (typeof ev.detail.value === 'string') {
                value = ev.detail.value
                emit('update', value)
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
</template>
