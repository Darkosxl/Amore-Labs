<script setup lang="ts">
import { ref } from 'vue'

const homeLink = ref('#/admin_console')
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

        <!-- Glass Card -->
        <div class="relative group">
          <!-- Glow effect -->
          <div class="absolute -inset-0.5 bg-gradient-to-r from-green-200 to-emerald-200 rounded-[2rem] blur opacity-40 group-hover:opacity-60 transition duration-500"></div>

          <div class="relative bg-white/60 backdrop-blur-xl ring-1 ring-white/50 border border-white/50 rounded-[2rem] p-8 sm:p-12 shadow-2xl text-center">

            <!-- Success Icon -->
            <div class="mx-auto flex h-20 w-20 items-center justify-center rounded-full bg-green-100 mb-6">
              <svg class="h-10 w-10 text-green-600" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" d="M4.5 12.75l6 6 9-13.5" />
              </svg>
            </div>

            <h2 class="text-3xl font-bold tracking-tighter text-gray-900 mb-4">
              Payment Successful
            </h2>

            <p class="text-base text-gray-600 mb-8">
              Thank you for your purchase. Your order has been processed successfully.
            </p>

            <div class="flex flex-col gap-4">
              <a :href="homeLink"
                 class="flex w-full justify-center rounded-full bg-black px-3 py-3.5 text-sm font-bold leading-6 text-white shadow-lg hover:bg-gray-800 hover:shadow-xl hover:-translate-y-0.5 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-black transition-all duration-200">
                Return to Dashboard
              </a>
            </div>

          </div>
        </div>

      </div>
    </main>

  </div>
</template>
