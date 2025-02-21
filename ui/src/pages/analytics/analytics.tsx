import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { AnalyticsConfigForm } from "@/components/forms/analytics-config-form"
import { BarChart3, LineChart, Users, Clock } from "lucide-react"

const stats = [
  {
    name: "Total Visitors",
    value: "24.7K",
    change: "+12%",
    icon: Users,
  },
  {
    name: "Page Views",
    value: "56.4K",
    change: "+8%",
    icon: BarChart3,
  },
  {
    name: "Avg. Session Duration",
    value: "2m 45s",
    change: "+23%",
    icon: Clock,
  },
  {
    name: "Bounce Rate",
    value: "24%",
    change: "-4%",
    icon: LineChart,
  },
]

export default function Analytics() {
  return (
    <div className="container mx-auto py-6">
      <div className="mb-8">
        <h1 className="text-2xl font-semibold text-gray-900">Analytics</h1>
        <p className="mt-2 text-sm text-gray-600">
          Track and analyze user behavior and engagement.
        </p>
      </div>

      {/* Stats Overview */}
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

      {/* Tabs */}
      <Tabs defaultValue="overview" className="space-y-6">
        <TabsList>
          <TabsTrigger value="overview">Overview</TabsTrigger>
          <TabsTrigger value="realtime">Real-time</TabsTrigger>
          <TabsTrigger value="behavior">Behavior</TabsTrigger>
          <TabsTrigger value="configuration">Configuration</TabsTrigger>
        </TabsList>

        <TabsContent value="overview">
          <div className="rounded-lg border bg-white p-6 shadow-sm">
            <div className="h-[400px] rounded-lg border-2 border-dashed border-gray-200 p-4">
              <p className="text-center text-gray-500">Analytics overview chart will be displayed here</p>
            </div>
          </div>
        </TabsContent>

        <TabsContent value="realtime">
          <div className="rounded-lg border bg-white p-6 shadow-sm">
            <div className="h-[400px] rounded-lg border-2 border-dashed border-gray-200 p-4">
              <p className="text-center text-gray-500">Real-time analytics will be displayed here</p>
            </div>
          </div>
        </TabsContent>

        <TabsContent value="behavior">
          <div className="rounded-lg border bg-white p-6 shadow-sm">
            <div className="h-[400px] rounded-lg border-2 border-dashed border-gray-200 p-4">
              <p className="text-center text-gray-500">User behavior analytics will be displayed here</p>
            </div>
          </div>
        </TabsContent>

        <TabsContent value="configuration">
          <div className="rounded-lg border bg-white p-6 shadow-sm">
            <AnalyticsConfigForm />
          </div>
        </TabsContent>
      </Tabs>
    </div>
  )
} 