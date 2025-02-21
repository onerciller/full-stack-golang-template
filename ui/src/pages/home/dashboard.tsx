import { Users, ShoppingCart, DollarSign, TrendingUp, ArrowRight, Bell } from "lucide-react"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"

const stats = [
  {
    name: "Total Users",
    value: "1,234",
    change: "+12%",
    icon: Users,
  },
  {
    name: "Total Orders",
    value: "456",
    change: "+8%",
    icon: ShoppingCart,
  },
  {
    name: "Revenue",
    value: "$12,345",
    change: "+23%",
    icon: DollarSign,
  },
  {
    name: "Growth",
    value: "89%",
    change: "+4%",
    icon: TrendingUp,
  },
]

const recentActivity = [
  {
    id: 1,
    type: "user",
    title: "New user registration",
    description: "John Doe registered a new account",
    timestamp: "2 hours ago",
    icon: Users,
  },
  {
    id: 2,
    type: "order",
    title: "New order received",
    description: "Order #12345 was placed",
    timestamp: "3 hours ago",
    icon: ShoppingCart,
  },
  {
    id: 3,
    type: "notification",
    title: "System update",
    description: "System maintenance completed",
    timestamp: "5 hours ago",
    icon: Bell,
  },
]

const topProducts = [
  { name: "Product A", sales: 120, revenue: "$1,200" },
  { name: "Product B", sales: 98, revenue: "$980" },
  { name: "Product C", sales: 75, revenue: "$750" },
  { name: "Product D", sales: 63, revenue: "$630" },
  { name: "Product E", sales: 45, revenue: "$450" },
]

export default function Dashboard() {
  return (
    <div className="container mx-auto py-6">
      <div className="mb-8">
        <h1 className="text-2xl font-semibold text-gray-900">Dashboard</h1>
        <p className="mt-2 text-sm text-gray-600">
          Welcome back! Here's an overview of your business.
        </p>
      </div>

      {/* Stats Grid */}
      <div className="mb-8 grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4">
        {stats.map((stat) => (
          <div
            key={stat.name}
            className="overflow-hidden rounded-lg bg-white px-4 py-5 shadow sm:p-6"
          >
            <div className="flex items-center">
              <div className="flex-shrink-0">
                <stat.icon className="h-6 w-6 text-gray-400" />
              </div>
              <div className="ml-5 w-0 flex-1">
                <p className="truncate text-sm font-medium text-gray-500">
                  {stat.name}
                </p>
                <p className="text-xl font-semibold text-gray-900">
                  {stat.value}
                </p>
              </div>
            </div>
            <div className="mt-4">
              <span
                className={`text-sm font-medium ${
                  stat.change.startsWith("+")
                    ? "text-green-600"
                    : "text-red-600"
                }`}
              >
                {stat.change}
              </span>
              <span className="text-sm text-gray-500"> from last month</span>
            </div>
          </div>
        ))}
      </div>

      {/* Main Content */}
      <div className="grid grid-cols-1 gap-6 lg:grid-cols-2">
        {/* Charts */}
        <div className="rounded-lg border bg-white p-6 shadow-sm">
          <div className="mb-4 flex items-center justify-between">
            <h2 className="text-lg font-medium text-gray-900">Revenue Overview</h2>
            <select className="rounded-md border-gray-300 text-sm">
              <option>Last 7 days</option>
              <option>Last 30 days</option>
              <option>Last 3 months</option>
            </select>
          </div>
          <div className="h-[300px] rounded-lg border-2 border-dashed border-gray-200 p-4">
            <p className="text-center text-gray-500">Revenue chart will be displayed here</p>
          </div>
        </div>

        {/* Recent Activity */}
        <div className="rounded-lg border bg-white p-6 shadow-sm">
          <div className="mb-4 flex items-center justify-between">
            <h2 className="text-lg font-medium text-gray-900">Recent Activity</h2>
            <button className="flex items-center text-sm text-primary hover:text-primary/80">
              View all
              <ArrowRight className="ml-1 h-4 w-4" />
            </button>
          </div>
          <div className="space-y-4">
            {recentActivity.map((activity) => (
              <div
                key={activity.id}
                className="flex items-start space-x-4 rounded-lg border p-4"
              >
                <div className="rounded-full bg-primary/10 p-2">
                  <activity.icon className="h-5 w-5 text-primary" />
                </div>
                <div className="flex-1 space-y-1">
                  <p className="text-sm font-medium text-gray-900">
                    {activity.title}
                  </p>
                  <p className="text-sm text-gray-500">{activity.description}</p>
                  <p className="text-xs text-gray-400">{activity.timestamp}</p>
                </div>
              </div>
            ))}
          </div>
        </div>

        {/* Top Products */}
        <div className="rounded-lg border bg-white p-6 shadow-sm">
          <h2 className="mb-4 text-lg font-medium text-gray-900">Top Products</h2>
          <div className="overflow-hidden">
            <table className="min-w-full">
              <thead>
                <tr className="border-b">
                  <th className="py-3 text-left text-sm font-semibold text-gray-900">Product</th>
                  <th className="py-3 text-right text-sm font-semibold text-gray-900">Sales</th>
                  <th className="py-3 text-right text-sm font-semibold text-gray-900">Revenue</th>
                </tr>
              </thead>
              <tbody>
                {topProducts.map((product, index) => (
                  <tr
                    key={product.name}
                    className={index !== topProducts.length - 1 ? "border-b" : ""}
                  >
                    <td className="py-3 text-sm text-gray-900">{product.name}</td>
                    <td className="py-3 text-right text-sm text-gray-500">{product.sales}</td>
                    <td className="py-3 text-right text-sm text-gray-900">{product.revenue}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>

        {/* Performance Metrics */}
        <div className="rounded-lg border bg-white p-6 shadow-sm">
          <h2 className="mb-4 text-lg font-medium text-gray-900">Performance Metrics</h2>
          <div className="space-y-4">
            <div>
              <div className="mb-1 flex items-center justify-between">
                <span className="text-sm font-medium text-gray-700">Conversion Rate</span>
                <span className="text-sm font-medium text-primary">3.6%</span>
              </div>
              <div className="h-2 rounded-full bg-gray-200">
                <div className="h-2 rounded-full bg-primary" style={{ width: "36%" }} />
              </div>
            </div>
            <div>
              <div className="mb-1 flex items-center justify-between">
                <span className="text-sm font-medium text-gray-700">Customer Satisfaction</span>
                <span className="text-sm font-medium text-primary">92%</span>
              </div>
              <div className="h-2 rounded-full bg-gray-200">
                <div className="h-2 rounded-full bg-primary" style={{ width: "92%" }} />
              </div>
            </div>
            <div>
              <div className="mb-1 flex items-center justify-between">
                <span className="text-sm font-medium text-gray-700">Average Order Value</span>
                <span className="text-sm font-medium text-primary">$86</span>
              </div>
              <div className="h-2 rounded-full bg-gray-200">
                <div className="h-2 rounded-full bg-primary" style={{ width: "72%" }} />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
} 