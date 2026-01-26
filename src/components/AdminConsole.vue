<script setup lang="ts">
import { ref, onMounted } from 'vue'
import ProductCard from './ProductCard.vue'

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

// User data
const user = ref({
  name: 'Loading...',
  email: '...',
  avatarGradient: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)'
})

// Loading state
const loadingSubscriptions = ref(true)

// Handle billing redirect
const handleBilling = async (productId: string) => {
  // Redirect to backend billing endpoint to create Stripe session
  const backendUrl = `http://localhost:8173/v1/billing/${productId}`

  // Create a form and submit it to redirect the user
  const form = document.createElement('form')
  form.method = 'POST'
  form.action = backendUrl
  document.body.appendChild(form)
  form.submit()
}

const handleSignOut = () => {
  // Clear access_token cookie
  document.cookie = 'access_token=; Max-Age=0; path=/; domain=' + window.location.hostname
  // Also clear for localhost if current domain is different, just in case
  document.cookie = 'access_token=; Max-Age=0; path=/;'

  // Clear any local storage if used
  localStorage.removeItem('user')

  // Redirect to Landing Page (root)
  window.location.hash = ''
}

// Product Data - will be updated with real subscription info
const products = ref([
  {
    id: 'rinova_ai',
    name: 'Rinova AI',
    description: 'Automated dispute resolution and revenue recovery.',
    status: 'inactive',
    renewalDate: null as string | null,
    trialDaysLeft: null as number | null,
    iconColor: 'text-blue-500',
    bgColor: 'bg-blue-500/20',
    borderColor: 'border-blue-500/30'
  },
  {
    id: 'voice_ai_italy_outbound',
    name: 'Outbound AI',
    description: 'Logistics precision, OTIF management, and carrier negotiation.',
    status: 'inactive',
    renewalDate: null as string | null,
    trialDaysLeft: null as number | null,
    iconColor: 'text-purple-500',
    bgColor: 'bg-purple-500/20',
    borderColor: 'border-purple-500/30'
  },
  {
    id: 'voice_ai_italy_inbound',
    name: 'Inbound AI',
    description: 'Seamless ERP integration and supply chain ingestion.',
    status: 'inactive',
    renewalDate: null as string | null,
    trialDaysLeft: null as number | null,
    iconColor: 'text-slate-500',
    bgColor: 'bg-slate-500/10',
    borderColor: 'border-slate-500/20'
  },
  {
    id: 'test_product',
    name: '🧪 Test Product',
    description: 'Internal testing only - visible to bscemarslan@gmail.com',
    status: 'inactive',
    renewalDate: null as string | null,
    trialDaysLeft: null as number | null,
    iconColor: 'text-yellow-500',
    bgColor: 'bg-yellow-500/20',
    borderColor: 'border-yellow-500/30'
  }
])

// Fetch subscriptions and update product statuses
const fetchSubscriptions = async () => {
  try {
    const response = await fetch('http://localhost:8173/v1/subscriptions', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include'
    })

    if (response.ok) {
      const data = await response.json()

      // Map subscriptions to products
      if (data.subscriptions && data.subscriptions.length > 0) {
        data.subscriptions.forEach((sub: { product_name: string; status: string; current_period_end: number }) => {
          const product = products.value.find(p => p.id === sub.product_name)
          if (product) {
            // Determine status based on subscription status
            if (sub.status === 'active') {
              product.status = 'active'
              // Format renewal date
              const renewalDate = new Date(sub.current_period_end * 1000)
              product.renewalDate = renewalDate.toLocaleDateString('en-US', {
                month: 'short',
                day: 'numeric',
                year: 'numeric'
              })
            } else if (sub.status === 'trialing') {
              product.status = 'trial'
              // Calculate days left in trial
              const trialEnd = new Date(sub.current_period_end * 1000)
              const now = new Date()
              const daysLeft = Math.ceil((trialEnd.getTime() - now.getTime()) / (1000 * 60 * 60 * 24))
              product.trialDaysLeft = daysLeft > 0 ? daysLeft : 0
            }
          }
        })
      }
    } else {
      console.warn('Failed to fetch subscriptions')
    }
  } catch (e) {
    console.error('Error fetching subscriptions:', e)
  } finally {
    loadingSubscriptions.value = false
  }
}

// Fetch User Data and Subscriptions
onMounted(async () => {
  // Check for payment failure in URL
  const urlParams = new URLSearchParams(window.location.search)
  if (urlParams.get('payment_failed') === 'true') {
    showNotification('Payment cancelled: You can try again anytime', 'error')
    // Clean up URL
    window.history.replaceState({}, '', window.location.pathname + window.location.hash.split('?')[0])
  }

  try {
    // Fetch user data
    const response = await fetch('http://localhost:8173/v1/me', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include'
    })

    if (response.ok) {
      const data = await response.json()
      user.value.email = data.email || user.value.email
      user.value.name = data.email ? data.email.split('@')[0] : 'Admin User'

      // Fetch subscriptions after user is loaded
      await fetchSubscriptions()
    } else {
      console.warn('Failed to fetch user data, using mock')
      if (response.status === 401) {
         window.location.hash = '#/signin'
      }
    }
  } catch (e) {
    console.error('Error connecting to backend:', e)
    user.value.name = 'Demo User'
    user.value.email = 'demo@amorelabs.com'
    loadingSubscriptions.value = false
  }
})
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
          <!-- Loading Skeleton -->
          <div v-if="loadingSubscriptions" class="space-y-6">
            <div v-for="i in 3" :key="i" class="relative rounded-[2rem] p-8 shadow-2xl border border-slate-800 overflow-hidden animate-pulse"
                 style="background: radial-gradient(circle at 0% 0%, #000000, #303030, #000000), linear-gradient(180deg, #505050, #000000); background-blend-mode: overlay, normal;">
              <div class="h-14 bg-slate-700/50 rounded-xl w-3/4"></div>
              <div class="h-4 bg-slate-700/30 rounded mt-4 w-1/2"></div>
            </div>
          </div>

          <!-- Product Cards -->
          <div v-else class="space-y-6">
            <!-- Rinova AI Card -->
            <div class="relative group" v-if="products[0]">
              <ProductCard :product="products[0]" @activate="handleBilling" />
            </div>

            <!-- Test Product Card (only for bscemarslan@gmail.com) -->
            <div class="relative group" v-if="products[3] && user.email === 'bscemarslan@gmail.com'">
              <ProductCard :product="products[3]" @activate="handleBilling" />
            </div>

            <!-- Outbound AI Card -->
            <div class="relative group" v-if="products[1]">
              <ProductCard :product="products[1]" @activate="handleBilling" />
            </div>

            <!-- Visual Connector with "Get Both" Button -->
            <div class="relative flex flex-col items-center justify-center -my-6 z-20">
              <!-- Vertical Connection Line -->
              <div class="h-12 w-0.5 bg-gradient-to-b from-slate-800 via-purple-500/50 to-slate-800"></div>

              <!-- Button in the middle -->
              <div class="relative group/btn my-[-4px]">
                <button
                  @click="handleBilling('voice_ai_italy_outbound-inbound')"
                  class="relative px-6 py-3 bg-slate-900 hover:bg-slate-800 text-white rounded-xl shadow-xl transition-all duration-300 border border-purple-500/30 hover:border-purple-500 hover:scale-105 group-hover/btn:shadow-purple-500/20"
                >
                  <div class="absolute inset-0 bg-gradient-to-r from-purple-500/10 to-pink-500/10 rounded-xl opacity-0 group-hover/btn:opacity-100 transition-opacity"></div>

                  <div class="relative flex items-center gap-3">
                    <!-- Icon -->
                    <div class="p-1.5 rounded-lg bg-gradient-to-br from-purple-500/20 to-pink-500/20 border border-purple-500/30">
                      <svg class="w-4 h-4 text-purple-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                      </svg>
                    </div>

                    <div class="text-left">
                      <div class="text-xs text-purple-300 font-medium">Bundle & Save</div>
                      <div class="text-sm font-bold text-white">Get Both Agents</div>
                    </div>

                    <!-- Chevron -->
                    <svg class="w-4 h-4 text-slate-500 group-hover/btn:text-white transition-colors ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
                    </svg>

                    <!-- Badge -->
                    <div class="absolute -top-3 -right-2 bg-gradient-to-r from-amber-400 to-orange-500 text-black text-[10px] font-black px-2 py-0.5 rounded-full shadow-lg transform rotate-3 group-hover/btn:rotate-6 transition-transform">
                      SAVE €€
                    </div>
                  </div>
                </button>
              </div>

              <!-- Vertical Connection Line -->
              <div class="h-12 w-0.5 bg-gradient-to-b from-slate-800 via-purple-500/50 to-slate-800"></div>
            </div>

            <!-- Inbound AI Card -->
            <div class="relative group" v-if="products[2]">
              <ProductCard :product="products[2]" @activate="handleBilling" />
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
      <div class="flex flex-col mr-2">
        <span class="text-sm font-bold text-gray-900 leading-tight">{{ user.name }}</span>
        <span class="text-xs text-gray-500 leading-tight">{{ user.email }}</span>
      </div>

      <!-- Sign Out Button -->
      <button @click.stop="handleSignOut" class="ml-2 p-2 rounded-full hover:bg-gray-100 text-gray-400 hover:text-red-500 transition-colors" title="Sign Out">
        <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
        </svg>
      </button>
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
