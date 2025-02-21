import { SettingsForm } from "@/components/forms/settings-form"

export default function Settings() {
  return (
    <div className="container mx-auto py-6">
      <div className="mb-8">
        <h1 className="text-2xl font-semibold text-gray-900">Settings</h1>
        <p className="mt-2 text-sm text-gray-600">
          Manage your application settings and preferences.
        </p>
      </div>

      <div className="rounded-lg border bg-white p-6 shadow-sm">
        <SettingsForm />
      </div>
    </div>
  )
} 