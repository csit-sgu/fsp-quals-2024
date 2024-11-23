import { createRouter, createWebHistory } from 'vue-router'
import MainPage from '../App.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: MainPage,
    },
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
  ],
})

export default router
