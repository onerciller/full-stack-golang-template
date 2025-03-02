import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { apiClient } from './client';

// Query key for users
export const usersKeys = {
  all: ['users'],
  lists: () => [...usersKeys.all, 'list'],
  list: (filters) => [...usersKeys.lists(), { filters }],
  details: () => [...usersKeys.all, 'detail'],
  detail: (id) => [...usersKeys.details(), id],
};

/**
 * Hook for fetching users with pagination and filters
 */
export function useUsers(params = {}) {
  const { page = 1, limit = 10, ...filters } = params;
  
  return useQuery({
    queryKey: usersKeys.list({ page, limit, ...filters }),
    queryFn: () => 
      apiClient(`/users?page=${page}&limit=${limit}${new URLSearchParams(filters)}`),
  });
}

/**
 * Hook for fetching a specific user by ID
 */
export function useUser(id) {
  return useQuery({
    queryKey: usersKeys.detail(id),
    queryFn: () => apiClient(`/users/${id}`),
    enabled: !!id, // Only run the query if we have an ID
  });
}

/**
 * Hook for creating a new user
 */
export function useCreateUser() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: (userData) => apiClient('/users', { method: 'POST', body: userData }),
    onSuccess: () => {
      // Invalidate and refetch users list
      queryClient.invalidateQueries({ queryKey: usersKeys.lists() });
    },
  });
}

/**
 * Hook for updating an existing user
 */
export function useUpdateUser() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: ({ id, ...userData }) => 
      apiClient(`/users/${id}`, { method: 'PUT', body: userData }),
    onSuccess: (data, variables) => {
      // Update the queries with the returned data
      queryClient.invalidateQueries({ queryKey: usersKeys.detail(variables.id) });
      queryClient.invalidateQueries({ queryKey: usersKeys.lists() });
    },
  });
}

/**
 * Hook for deleting a user
 */
export function useDeleteUser() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: (id) => apiClient(`/users/${id}`, { method: 'DELETE' }),
    onSuccess: (_, id) => {
      // Remove the user from the cache and refetch lists
      queryClient.removeQueries({ queryKey: usersKeys.detail(id) });
      queryClient.invalidateQueries({ queryKey: usersKeys.lists() });
    },
  });
} 