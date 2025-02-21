import { useMutation } from '@tanstack/react-query'

interface LoginRequest {
  username: string
  password: string
}

interface RegisterRequest {
  username: string
  email: string
  password: string
}

interface AuthResponse {
  access_token: string
}

interface AppError {
  code: string
  message: string
  status: number
}

const API_URL = 'http://localhost:3000'

export function useLogin() {
  return useMutation<AuthResponse, AppError, LoginRequest>({
    mutationFn: async (credentials) => {
      const response = await fetch(`${API_URL}/auth/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(credentials),
      })

      if (!response.ok) {
        const error = await response.json()
        throw error
      }

      return response.json()
    },
  })
}

export function useRegister() {
  return useMutation<AuthResponse, AppError, RegisterRequest>({
    mutationFn: async (userData) => {
      const response = await fetch(`${API_URL}/auth/register`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(userData),
      })

      if (!response.ok) {
        const error = await response.json()
        throw error
      }

      return response.json()
    },
  })
} 