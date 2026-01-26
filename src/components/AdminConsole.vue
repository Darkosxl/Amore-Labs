<script setup lang="ts">
import { ref, onMounted } from 'vue'

// Toast notification state
const showToast = ref(false)
const toastMessage = ref('')
const toastType = ref<'error' | 'success'>('error')

// Show toast notification
const showNotification = (message: string, type: 'error' | 'success' = 'error') => {
  toastMessage.value = message
  toastType.value = type
  showToast.value = true

  // Auto-hide after 4 seconds
  setTimeout(() => {
    showToast.value = false
  }, 4000)
}

// Mock data for UI demonstration
const user = ref({
  name: 'Loading...',
  email: '...',
  avatarGradient: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)'
})

// Fetch User Data
onMounted(async () => {
  // Check for payment failure in URL
  const urlParams = new URLSearchParams(window.location.search)
  if (urlParams.get('payment_failed') === 'true') {
    showNotification('Payment cancelled: You can try again anytime', 'error')
    // Clean up URL
    window.history.replaceState({}, '', window.location.pathname + window.location.hash.split('?')[0])
  }

  try {
    // Attempt to fetch real user data from backend
    // Note: This assumes the backend is running on port 8173 and CORS is configured
    const response = await fetch('http://localhost:8173/v1/me', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      // Important: Include cookies for authentication
      credentials: 'include'
    })

    if (response.ok) {
      const data = await response.json()
      user.value.email = data.email || user.value.email
      // Create a name from email if not provided (since backend only sends email/role)
      user.value.name = data.email ? data.email.split('@')[0] : 'Admin User'
    } else {
      console.warn('Failed to fetch user data, using mock')
      // Fallback or redirect to signin if 401
      if (response.status === 401) {
         window.location.hash = '#/signin'
      }
    }
  } catch (e) {
    console.error('Error connecting to backend:', e)
    user.value.name = 'Demo User'
    user.value.email = 'demo@amorelabs.com'
  }
})

const handleBilling = async (productId: string) => {
  // Redirect to backend billing endpoint to create Stripe session
  // formatting the URL to point to the backend
  const backendUrl = `http://localhost:8173/v1/billing/${productId}`

  // We create a form and submit it to redirect the user
  // This is often more reliable for redirects than fetch() if we expect a full page load
  // (which Stripe Checkout is)
  const form = document.createElement('form')
  form.method = 'POST'
  form.action = backendUrl
  document.body.appendChild(form)
  form.submit()
}

// Product Data
const products = ref([
  {
    id: 'rinova',
    name: 'Rinova AI',
    description: 'Automated dispute resolution and revenue recovery.',
    status: 'active',
    renewalDate: 'Feb 24, 2026',
    iconColor: 'text-blue-500',
    bgColor: 'bg-blue-500/20',
    borderColor: 'border-blue-500/30'
  },
  {
    id: 'voice_ai_italy_outbound',
    name: 'Outbound AI',
    description: 'Logistics precision, OTIF management, and carrier negotiation.',
    status: 'trial',
    trialDaysLeft: 14,
    iconColor: 'text-purple-500',
    bgColor: 'bg-purple-500/20',
    borderColor: 'border-purple-500/30'
  },
  {
    id: 'voice_ai_italy_inbound',
    name: 'Inbound AI',
    description: 'Seamless ERP integration and supply chain ingestion.',
    status: 'inactive',
    iconColor: 'text-slate-500',
    bgColor: 'bg-slate-500/10',
    borderColor: 'border-slate-500/20'
  }
])
</script>

<template>
  <div class="min-h-screen font-sans text-gray-900 bg-white relative overflow-auto">

    <!-- Background Effect -->
    <div class="absolute inset-0 opacity-20 pointer-events-none fixed"
         style="background-image: radial-gradient(#cbd5e1 1px, transparent 1px); background-size: 32px 32px;">
    </div>

    <main class="relative z-10 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-20 flex flex-col items-center">

      <div class="w-full max-w-4xl space-y-8">

        <!-- Header -->
        <div class="text-center mb-12">
          <h2 class="text-4xl font-bold tracking-tighter text-gray-900 mb-4">Admin Console</h2>
          <p class="text-gray-500">Manage your Amore Labs product suite.</p>
        </div>

        <!-- Product Cards Stack -->
        <div class="space-y-6">
          <div v-for="product in products" :key="product.id" class="relative group">
             <!-- Card Container (Dark Theme inspired by ManufacturingPowerSection) -->
             <div class="relative rounded-[2rem] p-8 shadow-2xl border border-slate-800 overflow-hidden transition-transform duration-300 hover:scale-[1.01]"
                  style="background: radial-gradient(circle at 0% 0%, #000000, #303030, #000000), linear-gradient(180deg, #505050, #000000); background-blend-mode: overlay, normal;">

                <!-- Background Grid -->
                <div class="absolute inset-0 opacity-20 pointer-events-none"
                     style="background-image: radial-gradient(#ffffff 1px, transparent 1px); background-size: 24px 24px;">
                </div>

                <!-- Content Grid -->
                <div class="relative z-10 grid grid-cols-1 md:grid-cols-12 gap-6 items-center">

                  <!-- Icon & Title -->
                  <div class="md:col-span-5 flex items-start space-x-5">
                    <div class="w-14 h-14 rounded-2xl flex-shrink-0 flex items-center justify-center border backdrop-blur-sm"
                         :class="[product.bgColor, product.borderColor]">
                       <!-- Simple Dynamic Icon based on name/color -->
                       <svg v-if="product.id === 'rinova'" class="w-7 h-7" :class="product.iconColor" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                           <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                       </svg>
                       <svg v-else-if="product.id === 'outbound'" class="w-7 h-7" :class="product.iconColor" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                           <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                       </svg>
                       <svg v-else class="w-7 h-7" :class="product.iconColor" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                           <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4m0 5c0 2.21-3.582 4-8 4s-8-1.79-8-4" />
                       </svg>
                    </div>
                    <div>
                      <h3 class="text-2xl font-bold text-white mb-1">{{ product.name }}</h3>
                      <p class="text-slate-400 text-sm leading-relaxed">{{ product.description }}</p>
                    </div>
                  </div>

                  <!-- Spacer/Divider for Mobile -->
                  <div class="hidden md:block md:col-span-1 border-r border-slate-700 h-12 mx-auto"></div>

                  <!-- Status -->
                  <div class="md:col-span-3">
                     <p class="text-xs text-slate-500 uppercase tracking-widest font-bold mb-2">Status</p>

                     <div v-if="product.status === 'active'" class="flex items-center space-x-2">
                        <span class="relative flex h-3 w-3">
                          <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-green-400 opacity-75"></span>
                          <span class="relative inline-flex rounded-full h-3 w-3 bg-green-500"></span>
                        </span>
                        <span class="text-white font-medium">Active Subscription</span>
                     </div>

                     <div v-else-if="product.status === 'trial'" class="flex items-center space-x-2">
                        <span class="relative flex h-3 w-3">
                          <span class="relative inline-flex rounded-full h-3 w-3 bg-purple-500"></span>
                        </span>
                        <span class="text-white font-medium">Free Trial</span>
                     </div>

                     <div v-else class="flex items-center space-x-2">
                        <span class="h-3 w-3 rounded-full bg-slate-600"></span>
                        <span class="text-slate-500 font-medium">Not Active</span>
                     </div>
                  </div>

                  <!-- Details/Action -->
                  <div class="md:col-span-3 text-right">
                      <div v-if="product.status === 'active'">
                        <p class="text-xs text-slate-500 mb-1">Renews on</p>
                        <p class="text-white font-mono text-sm">{{ product.renewalDate }}</p>
                        <button class="mt-3 text-xs text-slate-400 hover:text-white transition-colors">Manage Settings →</button>
                      </div>

                      <div v-else-if="product.status === 'trial'">
                        <p class="text-xs text-slate-500 mb-1">Time Remaining</p>
                        <p class="text-white font-bold text-lg">{{ product.trialDaysLeft }} Days</p>
                        <button @click="handleBilling(product.id)" class="mt-2 bg-purple-600 hover:bg-purple-500 text-white text-xs px-4 py-2 rounded-full transition-colors font-bold">
                          Upgrade Now
                        </button>
                      </div>

                      <div v-else>
                        <button @click="handleBilling(product.id)" class="border border-slate-600 hover:border-white text-slate-300 hover:text-white px-5 py-2 rounded-full text-sm transition-all">
                          Activate
                        </button>
                      </div>
                  </div>

                </div>
             </div>
          </div>
        </div>

      </div>

    </main>

    <!-- User Profile (Bottom Left) -->
    <div class="fixed bottom-6 left-6 z-50 flex items-center gap-4 bg-white/90 backdrop-blur-md px-4 py-3 rounded-full shadow-xl border border-white/50 hover:scale-105 transition-transform cursor-pointer">
      <!-- Gradient Circle Logo -->
      <div class="h-10 w-10 rounded-full flex-shrink-0 shadow-inner"
           :style="{ background: user.avatarGradient }">
      </div>

      <!-- User Info -->
      <div class="flex flex-col">
        <span class="text-sm font-bold text-gray-900 leading-tight">{{ user.name }}</span>
        <span class="text-xs text-gray-500 leading-tight">{{ user.email }}</span>
      </div>
    </div>

    <!-- Toast Notification -->
    <Transition
      enter-active-class="transition-all duration-300 ease-out"
      enter-from-class="opacity-0 translate-y-4"
      enter-to-class="opacity-100 translate-y-0"
      leave-active-class="transition-all duration-300 ease-in"
      leave-from-class="opacity-100 translate-y-0"
      leave-to-class="opacity-0 translate-y-4"
    >
      <div v-if="showToast"
           class="fixed top-6 right-6 z-50 max-w-md px-6 py-4 rounded-2xl shadow-2xl backdrop-blur-md border"
           :class="toastType === 'error' ? 'bg-red-500/90 border-red-400 text-white' : 'bg-green-500/90 border-green-400 text-white'">
        <div class="flex items-center gap-3">
          <!-- Icon -->
          <svg v-if="toastType === 'error'" class="w-6 h-6 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
          </svg>
          <svg v-else class="w-6 h-6 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <!-- Message -->
          <p class="font-medium">{{ toastMessage }}</p>
        </div>
      </div>
    </Transition>

  </div>
</template>
