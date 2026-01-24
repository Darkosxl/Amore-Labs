<script setup lang="ts">
import LandingPage from './components/LandingPage.vue'
import SignIn from './components/SignIn.vue'
import PaymentSuccess from './components/PaymentSuccess.vue'
import PaymentFailure from './components/PaymentFailure.vue'
import AdminConsole from './components/AdminConsole.vue'
import { ref, computed, onMounted } from 'vue'

const currentPath = ref(window.location.hash)

onMounted(() => {
  window.addEventListener('hashchange', () => {
    currentPath.value = window.location.hash
  })
})

const currentView = computed(() => {
  switch (currentPath.value) {
    case '#/signin':
      return SignIn
    case '#/success':
      return PaymentSuccess
    case '#/failure':
      return PaymentFailure
    case '#/admin_console':
      return AdminConsole
    default:
      return LandingPage
  }
})
</script>

<template>
  <component :is="currentView" />
</template>
