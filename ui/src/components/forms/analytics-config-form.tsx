import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import * as z from "zod"

import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Switch } from "@/components/ui/switch"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"

const formSchema = z.object({
  googleAnalyticsId: z.string().min(1, {
    message: "Google Analytics ID is required.",
  }),
  dataRetentionPeriod: z.string({
    required_error: "Please select a data retention period.",
  }),
  trackPageViews: z.boolean().default(true),
  trackClicks: z.boolean().default(true),
  trackFormSubmissions: z.boolean().default(true),
  trackErrors: z.boolean().default(true),
  samplingRate: z.string()
    .refine((val) => !isNaN(Number(val)) && Number(val) >= 0 && Number(val) <= 100, {
      message: "Sampling rate must be between 0 and 100",
    }),
  excludedPaths: z.string(),
  customDimensions: z.string(),
})

const retentionPeriods = [
  { value: "30", label: "30 days" },
  { value: "60", label: "60 days" },
  { value: "90", label: "90 days" },
  { value: "180", label: "180 days" },
  { value: "365", label: "1 year" },
]

export function AnalyticsConfigForm() {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      googleAnalyticsId: "",
      dataRetentionPeriod: "90",
      trackPageViews: true,
      trackClicks: true,
      trackFormSubmissions: true,
      trackErrors: true,
      samplingRate: "100",
      excludedPaths: "/admin/*, /api/*",
      customDimensions: "userType, userRole",
    },
  })

  function onSubmit(values: z.infer<typeof formSchema>) {
    console.log(values)
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <div className="space-y-8">
          {/* Basic Configuration */}
          <div>
            <h3 className="text-lg font-medium">Basic Configuration</h3>
            <div className="mt-4 space-y-6">
              <FormField
                control={form.control}
                name="googleAnalyticsId"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Google Analytics ID</FormLabel>
                    <FormControl>
                      <Input placeholder="G-XXXXXXXXXX" {...field} />
                    </FormControl>
                    <FormDescription>
                      Your Google Analytics 4 measurement ID.
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="dataRetentionPeriod"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Data Retention Period</FormLabel>
                    <Select onValueChange={field.onChange} defaultValue={field.value}>
                      <FormControl>
                        <SelectTrigger>
                          <SelectValue placeholder="Select retention period" />
                        </SelectTrigger>
                      </FormControl>
                      <SelectContent>
                        {retentionPeriods.map((period) => (
                          <SelectItem key={period.value} value={period.value}>
                            {period.label}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
                    <FormDescription>
                      How long to keep analytics data.
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>
          </div>

          {/* Tracking Options */}
          <div>
            <h3 className="text-lg font-medium">Tracking Options</h3>
            <div className="mt-4 space-y-4">
              <FormField
                control={form.control}
                name="trackPageViews"
                render={({ field }) => (
                  <FormItem className="flex flex-row items-center justify-between rounded-lg border p-4">
                    <div className="space-y-0.5">
                      <FormLabel className="text-base">Track Page Views</FormLabel>
                      <FormDescription>
                        Collect data about page views and navigation.
                      </FormDescription>
                    </div>
                    <FormControl>
                      <Switch
                        checked={field.value}
                        onCheckedChange={field.onChange}
                      />
                    </FormControl>
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="trackClicks"
                render={({ field }) => (
                  <FormItem className="flex flex-row items-center justify-between rounded-lg border p-4">
                    <div className="space-y-0.5">
                      <FormLabel className="text-base">Track Clicks</FormLabel>
                      <FormDescription>
                        Monitor user interactions with buttons and links.
                      </FormDescription>
                    </div>
                    <FormControl>
                      <Switch
                        checked={field.value}
                        onCheckedChange={field.onChange}
                      />
                    </FormControl>
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="trackFormSubmissions"
                render={({ field }) => (
                  <FormItem className="flex flex-row items-center justify-between rounded-lg border p-4">
                    <div className="space-y-0.5">
                      <FormLabel className="text-base">Track Form Submissions</FormLabel>
                      <FormDescription>
                        Collect data about form submissions and user input.
                      </FormDescription>
                    </div>
                    <FormControl>
                      <Switch
                        checked={field.value}
                        onCheckedChange={field.onChange}
                      />
                    </FormControl>
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="trackErrors"
                render={({ field }) => (
                  <FormItem className="flex flex-row items-center justify-between rounded-lg border p-4">
                    <div className="space-y-0.5">
                      <FormLabel className="text-base">Track Errors</FormLabel>
                      <FormDescription>
                        Monitor and collect error events and exceptions.
                      </FormDescription>
                    </div>
                    <FormControl>
                      <Switch
                        checked={field.value}
                        onCheckedChange={field.onChange}
                      />
                    </FormControl>
                  </FormItem>
                )}
              />
            </div>
          </div>

          {/* Advanced Settings */}
          <div>
            <h3 className="text-lg font-medium">Advanced Settings</h3>
            <div className="mt-4 space-y-6">
              <FormField
                control={form.control}
                name="samplingRate"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Sampling Rate (%)</FormLabel>
                    <FormControl>
                      <Input type="number" min="0" max="100" {...field} />
                    </FormControl>
                    <FormDescription>
                      Percentage of users to include in analytics (0-100).
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="excludedPaths"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Excluded Paths</FormLabel>
                    <FormControl>
                      <Input {...field} />
                    </FormControl>
                    <FormDescription>
                      Comma-separated list of paths to exclude from tracking (supports wildcards *).
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="customDimensions"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Custom Dimensions</FormLabel>
                    <FormControl>
                      <Input {...field} />
                    </FormControl>
                    <FormDescription>
                      Comma-separated list of custom dimensions to track.
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>
          </div>
        </div>

        <button
          type="submit"
          className="rounded-md bg-primary px-4 py-2 text-sm font-medium text-primary-foreground hover:bg-primary/90 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50"
        >
          Save Analytics Configuration
        </button>
      </form>
    </Form>
  )
} 