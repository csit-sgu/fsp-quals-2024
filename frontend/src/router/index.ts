import { createRouter, createWebHistory } from 'vue-router'
import MainPage from '@/MainPage.vue'
import ConfirmPage from '@/ConfirmPage.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/weekly',
      name: 'weekly',
      component: MainPage,
    },
    {
      path: '/table',
      name: 'table',
      component: MainPage,
    },
    {
      path: '/confirm/:confirm_id',
      name: 'confirm',
      component: ConfirmPage,
    },
    {
      path: '/',
      name: 'landing',
      component: MainPage,
    },
  ],
})

export default router
