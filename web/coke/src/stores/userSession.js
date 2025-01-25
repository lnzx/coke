import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useUserSession = defineStore('userSession', () => {
  const token = ref(localStorage.getItem('token') || '')

  const isLoggedIn = computed(() => !!token.value)

  // Action：设置 Token
  const setToken = (newToken) => {
    token.value = newToken
    localStorage.setItem('token', newToken) // 保存到 localStorage
  }

  // Action：清除 Token
  const clearToken = () => {
    token.value = ''
    localStorage.removeItem('token') // 从 localStorage 移除
  }

  return {
    token,
    isLoggedIn,
    setToken,
    clearToken,
  }
})
