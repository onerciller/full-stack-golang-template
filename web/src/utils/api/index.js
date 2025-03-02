// Export all API-related utilities
export * from './client';
export * from './users';
export * from './customers';

// Common hooks and utilities
import { useQuery } from '@tanstack/react-query';
import { apiClient } from './client';

/**
 * Generic hook for fetching data from any endpoint
 */
export function useFetch(endpoint, queryKey, options = {}) {
  return useQuery({
    queryKey: Array.isArray(queryKey) ? queryKey : [queryKey],
    queryFn: () => apiClient(endpoint),
    ...options,
  });
}

/**
 * Hook for fetching dashboard data
 */
export function useDashboardData() {
  return useQuery({
    queryKey: ['dashboard'],
    queryFn: () => apiClient('/dashboard'),
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
} 