<script setup lang="ts">
import { ref } from 'vue'

const homeLink = ref('#/')
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
          <div class="absolute -inset-0.5 bg-gradient-to-r from-red-200 to-orange-200 rounded-[2rem] blur opacity-40 group-hover:opacity-60 transition duration-500"></div>

          <div class="relative bg-white/60 backdrop-blur-xl ring-1 ring-white/50 border border-white/50 rounded-[2rem] p-8 sm:p-12 shadow-2xl text-center">

            <!-- Failure Icon -->
            <div class="mx-auto flex h-20 w-20 items-center justify-center rounded-full bg-red-100 mb-6">
              <svg class="h-10 w-10 text-red-600" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126zM12 15.75h.007v.008H12v-.008z" />
              </svg>
            </div>

            <h2 class="text-3xl font-bold tracking-tighter text-gray-900 mb-4">
              Payment Failed
            </h2>

            <p class="text-base text-gray-600 mb-8">
              We couldn't process your payment. Please try again or contact support if the issue persists.
            </p>

            <div class="flex flex-col gap-4">
              <button @click="window.location.reload()"
                      class="flex w-full justify-center rounded-full bg-black px-3 py-3.5 text-sm font-bold leading-6 text-white shadow-lg hover:bg-gray-800 hover:shadow-xl hover:-translate-y-0.5 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-black transition-all duration-200">
                Try Again
              </button>

              <a :href="homeLink" class="text-sm font-semibold text-gray-600 hover:text-gray-900 transition-colors">
                Return to Home
              </a>
            </div>

          </div>
        </div>

      </div>
    </main>

  </div>
</template>
