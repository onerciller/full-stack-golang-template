import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { Suspense, lazy } from "react"
import { AuthProvider } from '@/lib/context/auth-context'
import ProtectedRoute from '@/components/auth/protected-route'

// Layouts
const AdminLayout = lazy(() => import("./components/layouts/admin-layout"))

// Admin Pages
const Dashboard = lazy(() => import("@/pages/home/dashboard"))
const Users = lazy(() => import("@/pages/users/users"))
const Login = lazy(() => import("@/pages/login/login"))
const Settings = lazy(() => import("@/pages/settings/settings"))  
const Analytics = lazy(() => import("@/pages/analytics/analytics"))
const Profile = lazy(() => import("@/pages/profile/profile"))

const LoadingFallback = () => (
  <div className="flex h-screen items-center justify-center">
    <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary"></div>
  </div>
)

export default function App() {
  return (
    <Router>
      <AuthProvider>
        <Suspense fallback={<LoadingFallback />}>
          <Routes>
            {/* Auth Routes */}
            <Route path="/login" element={<Login />} />
            
            {/* Admin Routes */}
            <Route path="/" element={<AdminLayout />}>
              <Route index element={<ProtectedRoute><Dashboard /></ProtectedRoute>} />
              <Route path="users" element={<ProtectedRoute><Users /></ProtectedRoute>} />
              <Route path="settings" element={<ProtectedRoute><Settings /></ProtectedRoute>} />
              <Route path="analytics" element={<ProtectedRoute><Analytics /></ProtectedRoute>} />
              <Route path="profile" element={<ProtectedRoute><Profile /></ProtectedRoute>} />
            </Route>
            
          </Routes>
        </Suspense>
      </AuthProvider>
    </Router>
  )
}