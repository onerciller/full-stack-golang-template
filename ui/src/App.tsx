import { BrowserRouter, Routes, Route } from "react-router-dom"
import { Suspense, lazy } from "react"

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
    <BrowserRouter>
      <Suspense fallback={<LoadingFallback />}>
        <Routes>
          {/* Auth Routes */}
          <Route path="/login" element={<Login />} />
          
          {/* Admin Routes */}
          <Route path="/" element={<AdminLayout />}>
            <Route index element={<Dashboard />} />
            <Route path="users" element={<Users />} />
            <Route path="settings" element={<Settings />} />
            <Route path="analytics" element={<Analytics />} />
            <Route path="profile" element={<Profile />} />
          </Route>
          
        </Routes>
      </Suspense>
    </BrowserRouter>
  )
}