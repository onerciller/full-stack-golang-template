import { useQuery } from '@tanstack/react-query'

interface User {
  id: number
  username: string
  email: string
  created_at: string
  updated_at: string
  deleted_at: string | null
}

interface UsersResponse {
  users: User[]
}

interface AppError {
  code: string
  message: string
  status: number
}

const API_URL = 'http://localhost:3000'

export function useUsers() {
  return useQuery<UsersResponse, AppError>({
    queryKey: ['users'],
    queryFn: async () => {
      const token = localStorage.getItem('access_token')
      if (!token) {
        throw new Error('No access token found')
      }

      const response = await fetch(`${API_URL}/api/v1/users`, {
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      })

      if (!response.ok) {
        const error = await response.json()
        throw error
      }

      return response.json()
    },
  })
} 