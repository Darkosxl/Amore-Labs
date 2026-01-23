<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps({
  variant: {
    type: String,
    default: 'default'
  }
})

const variantClasses = computed(() => {
  switch (props.variant) {
    case 'dots':
      return 'bg-dots' // Placeholder for dots pattern if defined
    case 'glass':
      return 'bg-slate-800/50 backdrop-blur border-slate-700 text-slate-100'
    case 'gradient':
      return 'bg-gradient-to-br from-gray-800 to-gray-900 text-white'
    default:
      return 'bg-white text-slate-900'
  }
})
</script>

<template>
  <div
    :class="[
      'rounded-xl border shadow-sm',
      variantClasses
    ]"
  >
    <slot />
  </div>
</template>
