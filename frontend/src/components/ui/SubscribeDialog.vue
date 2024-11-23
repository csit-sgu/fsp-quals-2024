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
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { ref } from 'vue'
import { toast } from 'vue-sonner'

const props = defineProps<{
  searchFilters: string[]
}>()

const userEmail = ref('')
const updateEmail = (newValue: string | number) => {
  userEmail.value = newValue.toString()
}

const applyForSubscription = () => {
  // TODO(aguschin): backend needs to validate given email and return whether or not
  // subscription confirmation letter was sent successfully
  const promise = () =>
    new Promise<string>((resolve, reject) => {
      setTimeout(() => {
        console.log(userEmail.value)
        if (userEmail.value === 'test') {
          resolve(userEmail.value)
        } else {
          reject('Не удалось отправить письмо с подтверждением на указанную почту!')
        }
      }, 1500)
    })
  toast.promise(promise, {
    loading: 'Отправка подтверждения...',
    success: (email: string) => {
      return `Почта ${email} успешно подписана!`
    },
    error: (msg: string) => {
      return `Ошибка: ${msg}`
    },
  })
  return true
}
</script>

<template>
  <Dialog>
    <DialogTrigger as-child>
      <Button variant="outline">Подписаться на уведомления</Button>
    </DialogTrigger>
    <DialogContent
      @interact-outside="
        (event) => {
          const target = event.target as HTMLElement
          if (target?.closest('[data-sonner-toaster]')) return event.preventDefault()
        }
      "
    >
      <DialogHeader class="px-4 pt-3 pb-1">
        <DialogTitle class="text-2xl">Подписаться на уведомления</DialogTitle>
        <DialogDescription>
          Необходимо указать и подтвердить почту, чтобы начать получать уведомления для выбранных
          соревнований.
        </DialogDescription>
        <div class="flex flex-col">
          <div class="font-semibold">Выбранные параметры фильтров:</div>
          <!-- TODO(aguschin): we need better conditions for filters' existence -->
          <div
            v-if="props.searchFilters.filter((x) => x.trim()).length > 0"
            v-for="filterString in props.searchFilters"
          >
            {{ filterString }}
          </div>
          <div v-else>Не указано ни одного фильтра</div>

          <div class="font-semibold">Им соответствует {{ 1337 }} соревнований</div>
        </div>
      </DialogHeader>
      <DialogFooter class="flex flex-col sm:flex-col sm:gap-4 gap-4">
        <Input @update:model-value="updateEmail" type="email" placeholder="Ваш E-Mail..." />
        <Button @click="applyForSubscription">Подтвердить подписку</Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>
