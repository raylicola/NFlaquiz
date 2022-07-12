import { createRouter, createWebHistory } from 'vue-router'
import TopPage from '@/pages/TopPage'

const routes = [
    {
      path: '/',
      name: 'TopPage',
      component: TopPage
    },
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})
  
export default router