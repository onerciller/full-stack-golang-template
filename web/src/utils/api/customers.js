import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { apiClient } from './client';

// Query key for customers
export const customersKeys = {
  all: ['customers'],
  lists: () => [...customersKeys.all, 'list'],
  list: (filters) => [...customersKeys.lists(), { filters }],
  details: () => [...customersKeys.all, 'detail'],
  detail: (id) => [...customersKeys.details(), id],
};

/**
 * Hook for fetching customers with pagination and filters
 */
export function useCustomers(params = {}) {
  const { page = 1, limit = 10, ...filters } = params;
  
  return useQuery({
    queryKey: customersKeys.list({ page, limit, ...filters }),
    queryFn: () => 
      apiClient(`/customers?page=${page}&limit=${limit}${new URLSearchParams(filters)}`),
  });
}

/**
 * Hook for fetching a specific customer by ID
 */
export function useCustomer(id) {
  return useQuery({
    queryKey: customersKeys.detail(id),
    queryFn: () => apiClient(`/customers/${id}`),
    enabled: !!id, // Only run the query if we have an ID
  });
}

/**
 * Hook for creating a new customer
 */
export function useCreateCustomer() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: (customerData) => apiClient('/customers', { method: 'POST', body: customerData }),
    onSuccess: () => {
      // Invalidate and refetch customers list
      queryClient.invalidateQueries({ queryKey: customersKeys.lists() });
    },
  });
}

/**
 * Hook for updating an existing customer
 */
export function useUpdateCustomer() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: ({ id, ...customerData }) => 
      apiClient(`/customers/${id}`, { method: 'PUT', body: customerData }),
    onSuccess: (data, variables) => {
      // Update the queries with the returned data
      queryClient.invalidateQueries({ queryKey: customersKeys.detail(variables.id) });
      queryClient.invalidateQueries({ queryKey: customersKeys.lists() });
    },
  });
}

/**
 * Hook for deleting a customer
 */
export function useDeleteCustomer() {
  const queryClient = useQueryClient();
  
  return useMutation({
    mutationFn: (id) => apiClient(`/customers/${id}`, { method: 'DELETE' }),
    onSuccess: (_, id) => {
      // Remove the customer from the cache and refetch lists
      queryClient.removeQueries({ queryKey: customersKeys.detail(id) });
      queryClient.invalidateQueries({ queryKey: customersKeys.lists() });
    },
  });
} 