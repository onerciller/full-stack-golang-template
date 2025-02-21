import { Link, useLocation } from "react-router-dom"
import { cn } from "@/lib/utils"
import {
  LayoutDashboard,
  Users,
  Menu,
  X,
} from "lucide-react"
import { useState } from "react"

const navigation = [
  { name: "Dashboard", href: "/", icon: LayoutDashboard },
  { name: "Users", href: "/users", icon: Users },
]

const Logo = () => (
  <Link to="/" className="flex items-center gap-3">
    <div className="h-10 w-10 rounded-full bg-black flex items-center justify-center flex-shrink-0">
      <span className="text-white text-base font-bold">FGT</span>
    </div>
    <div>
      <h1 className="text-lg font-bold text-black leading-tight">
        Full Stack Golang
        <br />
        <span className="text-black/90">Template</span>
      </h1>
    </div>
  </Link>
)

export function Sidebar() {
  const [isSidebarOpen, setIsSidebarOpen] = useState(true)
  const location = useLocation()

  return (
    <>
      <div className="fixed top-0 left-0 z-40 lg:hidden">
        <button
          className="p-4 text-gray-600 hover:text-gray-900"
          onClick={() => setIsSidebarOpen(!isSidebarOpen)}
        >
          <Menu className="h-6 w-6" />
        </button>
      </div>

      <div
        className={cn(
          "fixed inset-y-0 left-0 z-50 w-72 transform bg-white shadow-lg transition-transform duration-300 ease-in-out lg:translate-x-0",
          isSidebarOpen ? "translate-x-0" : "-translate-x-full"
        )}
      >
        <div className="flex h-16 items-center justify-between border-b px-6">
          <Logo />
          <button
            className="lg:hidden p-2 text-gray-600 hover:text-gray-900 rounded-full hover:bg-gray-100"
            onClick={() => setIsSidebarOpen(false)}
          >
            <X className="h-5 w-5" />
          </button>
        </div>

        <nav className="mt-4 px-3">
          {navigation.map((item) => {
            const isActive = location.pathname === item.href
            return (
              <Link
                key={item.name}
                to={item.href}
                className={cn(
                  "flex items-center rounded-lg px-3 py-2.5 text-sm font-medium mb-0.5",
                  isActive
                    ? "bg-gray-100 text-gray-900"
                    : "text-gray-600 hover:bg-gray-50 hover:text-gray-900"
                )}
              >
                <item.icon
                  className={cn(
                    "mr-3 h-5 w-5",
                    isActive ? "text-gray-900" : "text-gray-400"
                  )}
                />
                {item.name}
              </Link>
            )
          })}
        </nav>
      </div>
    </>
  )
} 