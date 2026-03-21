<template>
  <div class="container">
    <h2>Регистрация</h2>
    <div v-if="error" class="error-text">{{ error }}</div>
    <form @submit.prevent="handleRegister">
      <input type="text" v-model="username" placeholder="Имя пользователя" required>
      <input type="password" v-model="password" placeholder="Пароль" required>
      <button type="submit" class="btn-receive">Зарегистрироваться</button>
    </form>
    <a @click="$router.push('/login')" class="nav-link">Уже есть аккаунт? Войти</a>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const username = ref('')
const password = ref('')
const error = ref('')
const router = useRouter()

const handleRegister = async () => {
  error.value = ''
  try {
    const res = await fetch('http://localhost:8080/api/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username: username.value, password: password.value })
    })
    const data = await res.json()
    if (res.ok) {
      router.push('/login')
    } else {
      error.value = data.error || 'Ошибка регистрации'
    }
  } catch (e) {
    error.value = 'Ошибка сети'
  }
}
</script>
