<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardFooter,
  CardTitle,
} from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { BACKEND_URL } from '@/lib/dataSource'
import axios from 'axios'
import type { AxiosResponse } from 'axios'
import { toast } from 'vue-sonner'
import { Toaster } from '@/components/ui/sonner'

const route = useRoute()

const confirmSubscription = () => {
  const confirmData = {
    confirmation: route.params.confirm_id,
  }
  console.log(confirmData)

  const tryConfirm = () =>
    new Promise<AxiosResponse>((resolve) => {
      resolve(axios.post(`${BACKEND_URL}/subscription/confirm`, confirmData))
    })

  toast.promise(tryConfirm, {
    loading: 'Ожидается подтверждение...',
    success: (_resp: AxiosResponse) => {
      return 'Ваша почта была успешно подтверждена!'
    },
    error: (resp: AxiosResponse) => {
      return `Произошла ошибка: ${resp.data.message || resp.toString()}`
    },
  })
  return true
}
</script>

<template>
  <Toaster richColors />
  <div class="w-dvw h-dvh flex place-content-center">
    <Card class="p-10 max-w-2xl sm:my-auto xs:min-h-dvh">
      <CardHeader class="pb-5">
        <CardTitle class="text-3xl text-center">Ваша подписка почти активирована!</CardTitle>
      </CardHeader>
      <CardContent class="flex flex-col text-lg">
        <div>
          Мы получили запрос на регистрацию Вашего адреса электронной почты в наши списки рассылки
          спортивных соревнований. Для подтверждения Вашей почты необходимо нажать на кнопку,
          расположенную ниже. Если запрос совершили не Вы, просто закройте эту страницу.
        </div>
      </CardContent>
      <CardFooter>
        <Button class="w-full text-lg p-6" @click="confirmSubscription">
          Подтвердить почтовый адрес
        </Button>
      </CardFooter>
    </Card>
  </div>
</template>
