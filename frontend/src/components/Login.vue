<template>
  <div class="container">
    <h2>Вход</h2>
    <div v-if="error" class="error-text">{{ error }}</div>
    <form @submit.prevent="handleLogin">
      <input type="text" v-model="username" placeholder="Имя пользователя" required>
      <input type="password" v-model="password" placeholder="Пароль" required>
      <button type="submit" class="btn-send">Войти</button>
    </form>
    <a @click="$router.push('/register')" class="nav-link">Нет аккаунта? Зарегистрироваться</a>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const username = ref('')
const password = ref('')
const error = ref('')
const router = useRouter()

const handleLogin = async () => {
  error.value = ''
  try {
    const res = await fetch('http://localhost:8080/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username: username.value, password: password.value })
    })
    const data = await res.json()
    if (res.ok) {
      localStorage.setItem('token', data.token)
      router.push('/')
    } else {
      error.value = data.error || 'Ошибка входа'
    }
  } catch (e) {
    error.value = 'Ошибка сети'
  }
}
</script>
