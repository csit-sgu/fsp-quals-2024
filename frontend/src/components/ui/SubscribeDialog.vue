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
import type { Condition, SubscriptionRequest } from '@/lib/dataSource'
import { BACKEND_URL } from '@/lib/dataSource'
import { ref } from 'vue'
import { toast } from 'vue-sonner'
import axios from 'axios'
import type { AxiosResponse } from 'axios'

const props = defineProps<{
  searchFilters: Condition
}>()

const userEmail = ref('')
const updateEmail = (newValue: string | number) => {
  userEmail.value = newValue.toString()
}

const applyForSubscription = () => {
  // TODO(aguschin): add the rest of filters
  const subscriptionData: SubscriptionRequest = {
    email: userEmail.value,
  }

  const trySubscribe = () =>
    new Promise<AxiosResponse>((resolve, reject) => {
      resolve(axios.post(`${BACKEND_URL}/subscription`, subscriptionData))
    })

  toast.promise(trySubscribe, {
    loading: 'Отправка письма с подтверждением...',
    success: (_resp: AxiosResponse) => {
      return `На Вашу почту отправлено письмо с подтверждением!`
    },
    error: (resp: AxiosResponse) => {
      return `Произошла ошибка: ${resp.data.message || resp.toString()}`
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
    <DialogContent @interact-outside="(event) => {
      const target = event.target as HTMLElement
      if (target?.closest('[data-sonner-toaster]')) return event.preventDefault()
    }
      ">
      <DialogHeader class="px-4 pt-3 pb-1">
        <DialogTitle class="text-2xl">Подписаться на уведомления</DialogTitle>
        <DialogDescription>
          Необходимо указать и подтвердить почту, чтобы начать получать уведомления для выбранных
          соревнований.
        </DialogDescription>
        <div class="flex flex-col">
          <div class="font-semibold">Выбранные параметры фильтров:</div>
          <!-- TODO(aguschin): we need better conditions for filters' existence -->
          <div v-if="props.searchFilters.title">
            Название соревнования: {{ props.searchFilters.title }}
          </div>
          <div v-if="props.searchFilters.sport">
            Вид спорта: {{ props.searchFilters.sport }}
          </div>
          <div v-if="props.searchFilters.country">
            Страна: {{ props.searchFilters.country }}
          </div>
          <div v-if="props.searchFilters.region">
            Регион: {{ props.searchFilters.region }}
          </div>
          <div v-if="props.searchFilters.locality">
            Населённый пункт: {{ props.searchFilters.locality }}
          </div>
          <div v-if="props.searchFilters.age">
            Возраст: {{ props.searchFilters.age }}
          </div>
          <div v-if="props.searchFilters.gender">
            Пол: {{ props.searchFilters.gender }}
          </div>
          <div v-if="props.searchFilters.event_type">
            Тип события: {{ props.searchFilters.event_type }}
          </div>
          <div v-if="props.searchFilters.event_scale">
            Уровень соревнования: {{ props.searchFilters.event_scale }}
          </div>
          <div v-if="props.searchFilters.additional_info && props.searchFilters.additional_info.length > 0">
            Дополнительная информация: {{ props.searchFilters.additional_info }}
          </div>

          <!-- <div v-else>Не указано ни одного фильтра</div> -->

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
