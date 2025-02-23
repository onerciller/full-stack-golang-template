import { useMutation, useQuery } from '@tanstack/react-query'
import { LoginRequest, RegisterRequest, AuthResponse, AppError, User } from '../model'


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

export function useAuthenticatedUser() {
  return useQuery<User, AppError>({
    queryKey: ['user'],
    queryFn: async () => {

      const token = localStorage.getItem('access_token')
      console.log("token", token)
      if (!token) {
        throw new Error('No access token found')
      }
      
      const response = await fetch(`${API_URL}/api/v1/user`, {
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      })

      console.log("response", response)

      if (!response.ok) {
        const error = await response.json()
        throw error
      }

      return response.json()
    },
  })
} 