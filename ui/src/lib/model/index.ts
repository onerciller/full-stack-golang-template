export interface LoginRequest {
    username: string
    password: string
  }
  
  export interface RegisterRequest {
    username: string
    email: string
    password: string
  }
  
  export interface AuthResponse {
    access_token: string
  }
  
  export interface AppError {
    code: string
    message: string
    status: number
  }


  export interface User {
    id: number
    username: string
    email: string
    created_at: string
    updated_at: string
    deleted_at: string | null
  }
  
  export interface UsersResponse {
    users: User[]
  }
  
  export interface AppError {
    code: string
    message: string
    status: number
  }