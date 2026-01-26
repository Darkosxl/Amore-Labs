<script setup lang="ts">
import { ref, onMounted } from 'vue'

// Get API URL from environment variable (falls back to localhost for dev)
const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8173'

const step = ref<'key' | 'request'>('key')
const masterKeyInput = ref('')
const isLoading = ref(false)
const error = ref('')

// Check if device is already authorized
onMounted(() => {
  const isAuthorized = localStorage.getItem('amore_device_authorized')
  if (isAuthorized === 'true') {
    isLoading.value = true
    window.location.href = `${API_URL}/auth/login`
  }
})

const verifyKey = async () => {
  if (!masterKeyInput.value) return

  isLoading.value = true
  error.value = ''

  try {
    const formData = new URLSearchParams()
    formData.append('key', masterKeyInput.value)

    const response = await fetch(`${API_URL}/auth/verify-masterkey`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      body: formData
    })

    if (!response.ok) {
       // Try to parse error message from JSON if possible
       try {
         const data = await response.json()
         throw new Error(data.error || 'Invalid Enterprise Key')
       } catch (e) {
         throw new Error('Invalid Enterprise Key')
       }
    }

    // Success
    localStorage.setItem('amore_device_authorized', 'true')
    window.location.href = `${API_URL}/auth/login`

  } catch (err: any) {
    console.error(err)
    error.value = err.message || 'Connection failed'
    isLoading.value = false
  }
}

const showRequestForm = () => {
  step.value = 'request'
}
</script>

<template>
  <div class="min-h-screen font-sans text-gray-900 bg-gradient-to-br from-white via-white to-slate-50">

    <main class="flex flex-col items-center justify-center min-h-screen px-4 sm:px-6 lg:px-8 relative">

      <!-- Logo (Top Left) -->
      <a href="#" class="absolute top-6 left-6 z-50">
          <img src="/amore-labs-logo.png" alt="Amore Labs" class="h-24 w-auto" />
      </a>

      <!-- Subtle Background Grid/Effect -->
      <div class="absolute inset-0 opacity-40 pointer-events-none"
           style="background-image: radial-gradient(#cbd5e1 1px, transparent 1px); background-size: 32px 32px;">
      </div>

      <div class="w-full max-w-md space-y-8 relative z-10">

        <!-- Header -->
        <div class="text-center">
          <h2 class="mt-6 text-4xl font-bold tracking-tighter text-gray-900">
            Enterprise Access
          </h2>
          <p class="mt-2 text-base text-gray-600">
            Secure Gateway for Authorized Personnel
          </p>
        </div>

        <!-- Glass Card Form -->
        <div class="relative group">
          <!-- Glow effect behind the card -->
          <div class="absolute -inset-0.5 bg-gradient-to-r from-gray-200 to-slate-200 rounded-[2rem] blur opacity-40 group-hover:opacity-60 transition duration-500"></div>

          <div class="relative bg-white/60 backdrop-blur-xl ring-1 ring-white/50 border border-white/50 rounded-[2rem] p-8 sm:p-12 shadow-2xl transition-all duration-300">

            <!-- STEP 1: ENTERPRISE KEY -->
            <form v-if="step === 'key'" @submit.prevent="verifyKey" class="space-y-6">
              <div>
                <label for="key" class="block text-sm font-bold leading-6 text-gray-900 ml-1">Enterprise Console Key</label>
                <div class="mt-2">
                  <input id="key" name="key" type="password" required v-model="masterKeyInput"
                         placeholder="••••••••••••••••"
                         class="block w-full rounded-xl border-0 py-3 px-4 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-200 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-black sm:text-sm sm:leading-6 bg-white/50 backdrop-blur-sm transition-all text-center tracking-widest" />
                </div>
                <p v-if="error" class="mt-2 text-sm text-red-600 font-medium text-center">{{ error }}</p>
              </div>

              <div class="pt-2 space-y-4">
                <button type="submit" :disabled="isLoading"
                        class="flex w-full justify-center rounded-full bg-black px-3 py-3.5 text-sm font-bold leading-6 text-white shadow-lg hover:bg-gray-800 hover:shadow-xl hover:-translate-y-0.5 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-black transition-all duration-200 disabled:opacity-70 disabled:cursor-not-allowed">
                  <span v-if="isLoading">Verifying & Redirecting...</span>
                  <span v-else>Access Console</span>
                </button>

                <button type="button" @click="showRequestForm" class="w-full text-center text-sm font-semibold text-gray-500 hover:text-gray-900 transition-colors">
                  Don't have a key? Request Access
                </button>
              </div>
            </form>

            <!-- STEP 2: REQUEST RECEIVED -->
            <div v-else-if="step === 'request'" class="text-center space-y-6">
               <div class="inline-flex items-center justify-center w-16 h-16 rounded-full bg-blue-100 mb-2">
                 <svg class="w-8 h-8 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                   <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
                 </svg>
               </div>

               <div>
                 <h3 class="text-lg font-bold text-gray-900">Request Received</h3>
                 <p class="mt-2 text-sm text-gray-600">
                   We have logged your request from this device. An administrator will review it shortly.
                 </p>
               </div>

               <button @click="step = 'key'" class="text-sm font-bold text-black hover:underline mt-4">
                 Back to Gateway
               </button>
            </div>

          </div>
        </div>

      </div>
    </main>

  </div>
</template>
