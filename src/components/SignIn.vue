<script setup lang="ts">


import { ref } from 'vue'

// Basic ref for plumbing later
const email = ref('')
const password = ref('')
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
            Welcome Back
          </h2>
          <p class="mt-2 text-base text-gray-600">
            Sign in to the Amore Labs Admin Console
          </p>
        </div>

        <!-- Glass Card Form -->
        <div class="relative group">
          <!-- Glow effect behind the card -->
          <div class="absolute -inset-0.5 bg-gradient-to-r from-gray-200 to-slate-200 rounded-[2rem] blur opacity-40 group-hover:opacity-60 transition duration-500"></div>

          <div class="relative bg-white/60 backdrop-blur-xl ring-1 ring-white/50 border border-white/50 rounded-[2rem] p-8 sm:p-12 shadow-2xl">
            <form class="space-y-6" @submit.prevent>

              <!-- Email Input -->
              <div>
                <label for="email" class="block text-sm font-bold leading-6 text-gray-900 ml-1">Email address</label>
                <div class="mt-2">
                  <input id="email" name="email" type="email" autocomplete="email" required v-model="email"
                         class="block w-full rounded-xl border-0 py-3 px-4 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-200 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-black sm:text-sm sm:leading-6 bg-white/50 backdrop-blur-sm transition-all" />
                </div>
              </div>

              <!-- Password Input -->
              <div>
                <div class="flex items-center justify-between ml-1">
                  <label for="password" class="block text-sm font-bold leading-6 text-gray-900">Password</label>
                  <div class="text-sm">
                    <a href="#" class="font-semibold text-gray-500 hover:text-gray-900 transition-colors">Forgot password?</a>
                  </div>
                </div>
                <div class="mt-2">
                  <input id="password" name="password" type="password" autocomplete="current-password" required v-model="password"
                         class="block w-full rounded-xl border-0 py-3 px-4 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-200 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-black sm:text-sm sm:leading-6 bg-white/50 backdrop-blur-sm transition-all" />
                </div>
              </div>

              <!-- Sign In Button -->
              <div class="pt-2">
                <button type="submit"
                        class="flex w-full justify-center rounded-full bg-black px-3 py-3.5 text-sm font-bold leading-6 text-white shadow-lg hover:bg-gray-800 hover:shadow-xl hover:-translate-y-0.5 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-black transition-all duration-200">
                  Sign in
                </button>
              </div>

            </form>
          </div>
        </div>

      </div>
    </main>

  </div>
</template>
