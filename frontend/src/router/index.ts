import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Messages from '../components/Messages.vue'
import Login from '../components/Login.vue'
import Register from '../components/Register.vue'

const routes: Array<RouteRecordRaw> = [
  { path: '/', component: Messages, meta: { requiresAuth: true } },
  { path: '/login', component: Login, meta: { requiresGuest: true } },
  { path: '/register', component: Register, meta: { requiresGuest: true } }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  if (to.meta.requiresAuth && !token) {
    next('/login')
  } else if (to.meta.requiresGuest && token) {
    next('/')
  } else {
    next()
  }
})

export default router
