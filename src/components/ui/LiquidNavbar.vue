<script setup lang="ts">
import { ref } from 'vue'

defineProps<{
  navItems?: { label: string; targetId: string }[]
}>()

defineEmits<{
  (e: 'openBooking'): void
}>()

const isOpen = ref(false)

const scrollTo = (id: string) => {
  isOpen.value = false
  const el = document.getElementById(id)
  if (el) {
    el.scrollIntoView({ behavior: 'smooth' })
  }
}
</script>

<template>
  <nav class="fixed top-6 left-1/2 -translate-x-1/2 z-50 w-[95%] max-w-6xl transition-all duration-300">
    <div class="relative rounded-full px-6 h-[72px] flex items-center justify-between backdrop-blur-lg bg-gradient-to-b from-white/40 to-white/5 border border-black ring-1 ring-white/50 shadow-[0_8px_32px_0_rgba(0,0,0,0.1),_inset_0_1px_1px_0_rgba(255,255,255,1),_inset_0_-2px_10px_0_rgba(0,0,0,0.05)]">

      <!-- Logo -->
      <div class="flex items-center space-x-3 cursor-pointer" @click="scrollTo('hero')">
        <img src="/amore-labs-logo.png" alt="Amore Labs" class="h-32 w-auto" />
      </div>

      <!-- Center Links (Desktop) -->
      <div class="hidden md:flex items-center space-x-8">
        <a v-for="item in navItems"
           :key="item.label"
           href="#"
           @click.prevent="scrollTo(item.targetId)"
           class="text-sm font-medium text-black/80 hover:text-black transition-colors"
        >
          {{ item.label }}
        </a>
      </div>

      <!-- CTA Button -->
      <button @click="$emit('openBooking')"
              class="hidden sm:block bg-black text-white px-6 py-2.5 rounded-full text-sm font-bold hover:bg-gray-800 transition-colors shadow-lg">
        Request Demo
      </button>

      <!-- Mobile Menu Button (Hamburger) -->
      <button class="md:hidden text-black p-2" @click="isOpen = !isOpen">
        <svg v-if="!isOpen" class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
        </svg>
        <svg v-else class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    </div>

    <!-- Mobile Menu Dropdown -->
    <div v-show="isOpen"
         class="absolute top-full left-0 mt-2 w-full rounded-2xl p-4 flex flex-col space-y-2 md:hidden backdrop-blur-lg bg-gradient-to-b from-white/40 to-white/5 border border-black ring-1 ring-white/50 shadow-[0_8px_32px_0_rgba(0,0,0,0.1),_inset_0_1px_1px_0_rgba(255,255,255,1),_inset_0_-2px_10px_0_rgba(0,0,0,0.05)]">
      <a v-for="item in navItems"
         :key="item.label"
         href="#"
         @click.prevent="scrollTo(item.targetId)"
         class="text-base font-medium text-black/80 hover:text-black px-4 py-2 hover:bg-black/5 rounded-lg transition-colors"
      >
        {{ item.label }}
      </a>
      <button @click="$emit('openBooking'); isOpen = false"
              class="mt-2 w-full bg-black text-white px-6 py-3 rounded-xl text-base font-bold hover:bg-gray-800 transition-colors text-center">
        Request Demo
      </button>
    </div>
  </nav>
</template>
