<template>
  <div class="container">
    <h2>Сообщения</h2>
    <button @click="logout" class="btn-delete" style="margin-bottom: 20px; width: 100%;">Выйти</button>
    <input v-model="text" type="text" placeholder="Введите текст сообщения">
    <div class="buttons">
      <button class="btn-send" @click="sendMessage('slow')">Отправить текст (5 сек)</button>
      <button class="btn-send-fast" @click="sendMessage('fast')">Отправить текст (2 сек)</button>
      <button class="btn-receive" @click="fetchMessages">Получить сообщения</button>
    </div>
    <ul class="message-list">
      <li class="message-item" v-for="msg in messages" :key="msg.id">
        <span>{{ msg.text }}</span>
        <button class="btn-delete" @click="deleteMessage(msg.id)">Удалить</button>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

interface Message {
  id: number;
  text: string;
}

const text = ref('')
const messages = ref<Message[]>([])
const router = useRouter()

const getAuthHeaders = (): Record<string, string> => {
  return {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${localStorage.getItem('token')}`
  }
}

const handleAuthError = (res: Response): boolean => {
  if (res.status === 401) {
    logout()
    return true
  }
  return false
}

const logout = (): void => {
  localStorage.removeItem('token')
  router.push('/login')
}

const sendMessage = async (type: string): Promise<void> => {
  if (!text.value.trim()) return;
  try {
    const res = await fetch('http://localhost:8080/api/messages', {
      method: 'POST',
      headers: getAuthHeaders(),
      body: JSON.stringify({ text: text.value, type: type })
    });
    if (handleAuthError(res)) return;
    if (res.ok) {
      text.value = '';
      fetchMessages();
    } else {
      const data = await res.json();
      alert(`Error sending message: ${data.error || res.statusText}`);
    }
  } catch (e) {
    alert(`Error sending message: ${e instanceof Error ? e.message : String(e)}`);
  }
}

const fetchMessages = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/messages', {
      headers: getAuthHeaders()
    });
    if (handleAuthError(res)) return;
    const data = await res.json();
    messages.value = data || [];
  } catch (e) {
    alert(`Error fetching messages: ${e instanceof Error ? e.message : String(e)}`);
  }
}

const deleteMessage = async (id: number): Promise<void> => {
  try {
    const res = await fetch(`http://localhost:8080/api/messages/${id}`, {
      method: 'DELETE',
      headers: getAuthHeaders()
    });
    if (handleAuthError(res)) return;
    if (res.ok) {
      messages.value = messages.value.filter(m => m.id !== id);
    } else {
      const data = await res.json();
      alert(`Error deleting message: ${data.error || res.statusText}`);
    }
  } catch (e) {
    alert(`Error deleting message: ${e instanceof Error ? e.message : String(e)}`);
  }
}

onMounted(() => {
  fetchMessages()
})
</script>
