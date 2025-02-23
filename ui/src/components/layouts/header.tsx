import { Bell, User, LogOut } from "lucide-react"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Link } from "react-router-dom"
import { useAuth } from "@/lib/context/auth-context"
import { useAuthenticatedUser } from '@/lib/queries/auth'

const MobileLogo = () => (
  <Link to="/" className="flex items-center gap-2">
    <div className="h-8 w-8 rounded-full bg-black flex items-center justify-center flex-shrink-0">
      <span className="text-white text-sm font-bold">FGT</span>
    </div>
    <h1 className="text-sm font-bold text-black">
      Full Stack Golang
      <span className="text-black/90 ml-1">Template</span>
    </h1>
  </Link>
)

export function Header() {
  const { isAuthenticated, logout } = useAuth()
  const { data: user } = useAuthenticatedUser()
  console.log(user)
  return (
    <header className="sticky top-0 z-40 border-b bg-white">
      <div className="flex h-16 items-center justify-between px-4 sm:px-6 lg:px-8">
        <div className="lg:hidden">
          <MobileLogo />
        </div>
        
        <div className="flex items-center gap-3 ml-auto">
          <button className="rounded-full p-2 text-gray-400 hover:bg-gray-100 hover:text-gray-500 transition-colors">
            <Bell className="h-5 w-5" />
          </button>
          
          {isAuthenticated && (
            <DropdownMenu>
              <DropdownMenuTrigger className="flex items-center gap-3 outline-none">
                <div className="hidden sm:flex sm:flex-col sm:items-end sm:justify-center">
                <p className="text-sm font-medium text-gray-900">{user?.username}</p>
                <p className="text-xs text-gray-500">{user?.email}</p>
              </div>
              <div className="rounded-full bg-gray-100 p-2 text-gray-600 hover:bg-gray-200 transition-colors">
                <User className="h-5 w-5" />
              </div>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end" className="w-56">
              <DropdownMenuLabel>My Account</DropdownMenuLabel>
              <DropdownMenuSeparator />                          
              <DropdownMenuItem className="text-red-600 cursor-pointer" onClick={logout}>
                <LogOut className="mr-2 h-4 w-4" />
                <span>Log out</span>
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
          )}
        </div>
      </div>
    </header>
  )
} 