import axios from 'axios'
import { useUserSession } from '@/stores/userSession'

let api

export function createApi() {
  // Here we set the base URL for all requests made to the api
  api = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
  })

  // We set an interceptor for each request to include Bearer token to the request if user is logged in
  api.interceptors.request.use((config) => {
    const userSession = useUserSession()

    if (userSession.isLoggedIn) {
      config.headers = {
        ...config.headers, // 保留已有的 headers
        Authorization: `${userSession.token}`, // 添加 Token
      }
    }
    return config
  })

  return api
}

export function useApi() {
  if (!api) {
    createApi()
  }
  return api
}
