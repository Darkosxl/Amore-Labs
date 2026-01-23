<script setup lang="ts">
const navItems = [
  'Why Adversarial Agents',
  'The Metrics',
  'Negotiation Engine',
  'Integrations'
]

import IntegrationSection from './IntegrationSection.vue'
import IntegrationsGrid from './IntegrationsGrid.vue'
import BookingModal from './BookingModal.vue'
import ManufacturingPowerSection from './ManufacturingPowerSection.vue' // New Section
import { ref } from 'vue'

const showBooking = ref(false)

const failureCards = [
  {
    title: 'Customer Order Fulfillment',
    subtitle: 'The Internal Check',
    definition: 'Did we have the cases to put on the truck?',
    target: '95%+',
    pain: 'Production issues or raw material shortages lead to immediate relationship trouble.'
  },
  {
    title: 'OTIF',
    subtitle: 'The Handshake',
    definition: 'Logistics precision. Arriving strictly on time and in full.',
    target: '98%',
    pain: 'Arriving early clogs backrooms; arriving late causes empty shelves.'
  },
  {
    title: 'In-Stock Measures',
    subtitle: 'The Consumer View',
    definition: 'Is the product on the shelf when the shopper reaches for it?',
    target: '90-95%',
    pain: 'Hardest to control due to "phantom inventory." If this fails, customers walk away.'
  }
]
</script>

<template>
  <div class="min-h-screen font-sans text-gray-900 bg-white">
    <!-- Navbar -->
    <header class="fixed top-0 w-full z-50 bg-white/80 backdrop-blur-sm border-b border-gray-100">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="relative flex justify-between items-center h-20">
          <!-- Left: Nav Links -->
          <nav class="hidden md:flex space-x-8 z-10">
            <a
              v-for="item in navItems"
              :key="item"
              href="#"
              class="text-[11px] font-medium text-gray-500 hover:text-gray-900 tracking-wider uppercase"
            >
              {{ item }}
            </a>
          </nav>

          <!-- Center: Logo -->
          <div class="absolute left-1/2 transform -translate-x-1/2 flex items-center">
             <img src="/amore-labs-logo.png" alt="Amore Labs" class="h-12 w-auto" />
          </div>

          <!-- Right: Button -->
          <div class="flex items-center space-x-4 z-10">
            <button @click="showBooking = true" class="bg-black text-white px-5 py-2.5 rounded-md text-sm font-medium hover:bg-gray-800 transition-colors">
              Request Demo
            </button>
          </div>
        </div>
      </div>
    </header>

    <!-- Hero Section -->
    <main class="pt-32 pb-16 px-4 sm:px-6 lg:px-8 max-w-7xl mx-auto">
      <div class="text-center mb-24">

        <h1 class="text-4xl sm:text-5xl md:text-6xl font-bold tracking-tight text-gray-00 mb-6">
         Adversarial Coordinator Agents<br class="hidden sm:block" />
        </h1>

        <p class="text-lg sm:text-xl text-gray-600 max-w-3xl mx-auto mb-10">
          Don't just accept the fine. When production lines break, our Adversarial Agent negotiates substitutions, prevents on-shelf gaps, and secures Best-in-Class status before the truck leaves.
        </p>

        <div class="flex justify-center">
          <button @click="showBooking = true" class="bg-gray-900 text-white px-8 py-3 rounded-md text-base font-medium hover:bg-gray-800 transition-colors">
            Talk to Founders
          </button>
        </div>
      </div>

      <!-- Integration Section (Recreated from Image) -->
      <IntegrationSection />

      <!-- Integrations Grid (Octagonal Logos) -->
      <IntegrationsGrid @openBooking="showBooking = true" />

      <!-- Manufacturing Power Section (New - moved below) -->
      <ManufacturingPowerSection />

      <!-- Live Action Workflow (Redesigned) -->
      <section class="py-24 bg-white border-t border-gray-100">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div class="text-center mb-24">
            <h2 class="text-4xl font-bold text-gray-900 mb-6">How Amore Labs Transforms Supply Chain</h2>
            <p class="text-lg text-gray-600 max-w-2xl mx-auto">Enterprise-grade automation with zero training required.</p>
          </div>

          <div class="max-w-6xl mx-auto">
            <!-- Grid Container -->
            <div class="grid grid-cols-1 md:grid-cols-3 gap-y-24 gap-x-12 relative">

              <!-- STEP 1 -->
              <div class="flex flex-col items-center text-center relative group">
                <!-- Arrow 1 -> 2 (Right) -->
                <div class="hidden md:block absolute top-10 -right-[4.5rem] w-16 text-gray-300">
                  <svg viewBox="0 0 24 24" fill="none" class="w-full h-6" stroke="currentColor" stroke-width="2">
                     <path d="M5 12h14M15 8l4 4-4 4" stroke-linecap="round" stroke-linejoin="round"/>
                  </svg>
                </div>

                <div class="w-20 h-20 bg-blue-50 rounded-full flex items-center justify-center mb-6 shadow-sm ring-1 ring-blue-100">
                  <svg class="w-8 h-8 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                  </svg>
                </div>
                <div class="text-xs font-bold text-gray-400 uppercase tracking-widest mb-2">Step 1</div>
                <h3 class="text-xl font-bold text-gray-900 mb-3">The Trigger</h3>
                <p class="text-sm text-gray-600 leading-relaxed max-w-xs">D&S reviews retailer In-Stock measures and detects stockouts.</p>
              </div>

              <!-- STEP 2 -->
              <div class="flex flex-col items-center text-center relative group">
                <!-- Arrow 2 -> 3 (Right) -->
                 <div class="hidden md:block absolute top-10 -right-[4.5rem] w-16 text-gray-300">
                  <svg viewBox="0 0 24 24" fill="none" class="w-full h-6" stroke="currentColor" stroke-width="2">
                     <path d="M5 12h14M15 8l4 4-4 4" stroke-linecap="round" stroke-linejoin="round"/>
                  </svg>
                </div>

                <div class="w-20 h-20 bg-purple-50 rounded-full flex items-center justify-center mb-6 shadow-sm ring-1 ring-purple-100">
                  <svg class="w-8 h-8 text-purple-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
                  </svg>
                </div>
                <div class="text-xs font-bold text-gray-400 uppercase tracking-widest mb-2">Step 2</div>
                <h3 class="text-xl font-bold text-gray-900 mb-3">Referral</h3>
                <p class="text-sm text-gray-600 leading-relaxed max-w-xs">Route requirements to CSC Adversarial Agents.</p>
              </div>

              <!-- STEP 3 -->
              <div class="flex flex-col items-center text-center relative group">
                <!-- Arrow 3 -> 4 (Down) -->
                <div class="hidden md:block absolute -bottom-16 left-1/2 -translate-x-1/2 h-16 text-gray-300">
                   <svg viewBox="0 0 24 24" fill="none" class="h-full w-6" stroke="currentColor" stroke-width="2">
                     <path d="M12 5v14M8 15l4 4 4-4" stroke-linecap="round" stroke-linejoin="round"/>
                  </svg>
                </div>

                <div class="w-20 h-20 bg-orange-50 rounded-full flex items-center justify-center mb-6 shadow-sm ring-1 ring-orange-100">
                  <svg class="w-8 h-8 text-orange-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 7h12m0 0l-4-4m4 4l-4 4m0 6H4m0 0l4 4m-4-4l4-4" />
                  </svg>
                </div>
                <div class="text-xs font-bold text-gray-400 uppercase tracking-widest mb-2">Step 3</div>
                <h3 class="text-xl font-bold text-gray-900 mb-3">The Pitch</h3>
                <p class="text-sm text-gray-600 leading-relaxed max-w-xs">Adversarial Agent proposes offers to maximize fulfillment.</p>
              </div>

              <!-- STEP 6 (Bottom Left) -->
              <!-- Visual Order in DOM matches CSS Grid, but visually reversed by arrows -->
              <div class="flex flex-col items-center text-center relative group order-last md:order-none">

                <div class="w-20 h-20 bg-indigo-50 rounded-full flex items-center justify-center mb-6 shadow-sm ring-1 ring-indigo-100">
                   <svg class="w-8 h-8 text-indigo-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                     <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 002 2h2a2 2 0 002-2z" />
                   </svg>
                </div>
                <div class="text-xs font-bold text-gray-400 uppercase tracking-widest mb-2">Step 6</div>
                <h3 class="text-xl font-bold text-gray-900 mb-3">Visibility</h3>
                <p class="text-sm text-gray-600 leading-relaxed max-w-xs">Real-time sync of Reason Codes and delivery status.</p>
              </div>

              <!-- STEP 5 (Bottom Middle) -->
              <div class="flex flex-col items-center text-center relative group order-5 md:order-none">
                <!-- Arrow 5 -> 6 (Left) -->
                <div class="hidden md:block absolute top-10 -left-[4.5rem] w-16 text-gray-300">
                  <svg viewBox="0 0 24 24" fill="none" class="w-full h-6" stroke="currentColor" stroke-width="2">
                     <path d="M19 12H5M9 8l-4 4 4 4" stroke-linecap="round" stroke-linejoin="round"/>
                  </svg>
                </div>

                <div class="w-20 h-20 bg-green-50 rounded-full flex items-center justify-center mb-6 shadow-sm ring-1 ring-green-100">
                  <svg class="w-8 h-8 text-green-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                  </svg>
                </div>
                <div class="text-xs font-bold text-gray-400 uppercase tracking-widest mb-2">Step 5</div>
                <h3 class="text-xl font-bold text-gray-900 mb-3">Execution</h3>
                <p class="text-sm text-gray-600 leading-relaxed max-w-xs">Automated RFX & PO generation.</p>
              </div>

              <!-- STEP 4 (Bottom Right) -->
              <div class="flex flex-col items-center text-center relative group order-4 md:order-none">
                <!-- Arrow 4 -> 5 (Left) -->
                <div class="hidden md:block absolute top-10 -left-[4.5rem] w-16 text-gray-300">
                  <svg viewBox="0 0 24 24" fill="none" class="w-full h-6" stroke="currentColor" stroke-width="2">
                     <path d="M19 12H5M9 8l-4 4 4 4" stroke-linecap="round" stroke-linejoin="round"/>
                  </svg>
                </div>

                <div class="w-20 h-20 bg-yellow-50 rounded-full flex items-center justify-center mb-6 shadow-sm ring-1 ring-yellow-100">
                   <svg class="w-8 h-8 text-yellow-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                     <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                   </svg>
                </div>
                <div class="text-xs font-bold text-gray-400 uppercase tracking-widest mb-2">Step 4</div>
                <h3 class="text-xl font-bold text-gray-900 mb-3">The Counter</h3>
                <p class="text-sm text-gray-600 leading-relaxed max-w-xs">Retailer responses reviewed by CSC.</p>
              </div>

            </div>
          </div>
        </div>
      </section>

      <!-- Three Views of Failure -->
      <section class="py-20 border-t border-gray-100">
        <div class="text-center mb-16">
          <h2 class="text-3xl font-bold text-gray-900 mb-4">Three Views of Failure</h2>
          <p class="text-gray-600">Why the Adversarial Agent is a strategic necessity.</p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div v-for="(card, index) in failureCards" :key="index" class="p-8 rounded-lg border border-gray-100 bg-gray-50 hover:shadow-md transition-shadow">
            <div class="flex justify-between items-start mb-6">
              <div>
                <h3 class="text-lg font-bold text-gray-900">{{ card.title }}</h3>
                <span class="text-xs text-gray-500 uppercase tracking-wider">{{ card.subtitle }}</span>
              </div>
              <div class="text-right">
                <div class="text-sm text-gray-400">Target</div>
                <div class="text-xl font-mono font-bold text-gray-900">{{ card.target }}</div>
              </div>
            </div>

            <p class="text-sm text-gray-800 font-medium mb-4">"{{ card.definition }}"</p>
            <div>
              <span class="text-xs text-red-600 font-bold uppercase block mb-1">Risk Factor</span>
              <p class="text-sm text-gray-600">{{ card.pain }}</p>
            </div>
          </div>
        </div>
      </section>

      <!-- Feature Deep Dive: Coffee Scenario (Clean UI Version) -->
      <section class="py-20 border-t border-gray-100">
        <h2 class="text-3xl font-bold text-center text-gray-900 mb-16">Interactive Scenario</h2>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 max-w-5xl mx-auto">
          <!-- Left: Alert Card -->
          <div class="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
            <div class="bg-red-50 p-4 border-b border-red-100 flex items-center space-x-2">
              <span class="w-2 h-2 bg-red-500 rounded-full"></span>
              <span class="text-xs font-bold text-red-700 uppercase">System Alert: Critical</span>
            </div>
            <div class="p-8">
              <div class="space-y-6">
                <div>
                  <div class="text-xs text-gray-400 uppercase tracking-wide">SKU</div>
                  <div class="text-lg font-bold text-gray-900">500g Coffee Jar</div>
                </div>
                <div>
                  <div class="text-xs text-gray-400 uppercase tracking-wide">Status</div>
                  <div class="text-red-600 font-bold">Stockout Detected</div>
                </div>
                <div class="grid grid-cols-2 gap-4">
                  <div class="p-4 bg-gray-50 rounded">
                    <div class="text-xs text-gray-500">Demand</div>
                    <div class="font-bold text-gray-900">1,000 Cases</div>
                  </div>
                  <div class="p-4 bg-red-50 rounded">
                    <div class="text-xs text-red-500">Available</div>
                    <div class="font-bold text-red-700">100 Cases</div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Right: Chat Interface -->
          <div class="bg-gray-50 rounded-xl border border-gray-200 p-6 flex flex-col h-full">
            <div class="flex-1 space-y-4">
               <div class="flex items-start space-x-3">
                 <div class="w-8 h-8 rounded-full bg-blue-100 flex items-center justify-center text-blue-600 text-xs font-bold">AA</div>
                 <div class="bg-white p-3 rounded-tr-xl rounded-br-xl rounded-bl-xl shadow-sm text-sm text-gray-700">
                   Connecting to Retailer Buyer...
                 </div>
               </div>

               <div class="flex items-start space-x-3">
                 <div class="w-8 h-8 rounded-full bg-blue-100 flex items-center justify-center text-blue-600 text-xs font-bold">AA</div>
                 <div class="bg-white p-3 rounded-tr-xl rounded-br-xl rounded-bl-xl shadow-sm text-sm text-gray-700">
                   We are short on 500g jars. To prevent empty shelves, we propose substituting with 500 cases of 1kg Family Packs.
                 </div>
               </div>

               <div class="flex justify-center my-4">
                 <span class="text-xs text-gray-400 italic">Strategy: Maintaining stock coverage.</span>
               </div>

               <div class="flex justify-center">
                 <div class="bg-green-50 text-green-700 px-4 py-2 rounded-full text-xs font-bold border border-green-200">
                   AGREED. ORDER MODIFIED.
                 </div>
               </div>
            </div>
          </div>
        </div>
      </section>

      <!-- Why Negotiation Matters -->
      <section class="py-20 border-t border-gray-100">
        <h2 class="text-3xl font-bold text-center text-gray-900 mb-12">Visibility Over Perfection</h2>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-12">
          <div>
            <h3 class="text-lg font-bold text-gray-900 mb-3">Protecting the Scorecard</h3>
            <p class="text-gray-600 text-sm leading-relaxed">Finding a solution is a strategic necessity. Fulfilling value/volume via substitution maintains the relationship score.</p>
          </div>
          <div>
            <h3 class="text-lg font-bold text-gray-900 mb-3">The "Best-in-Class" Distinction</h3>
            <p class="text-gray-600 text-sm leading-relaxed">Proactive communication separates top-tier partners from the rest. The retailer demands "No Surprises."</p>
          </div>
          <div>
            <h3 class="text-lg font-bold text-gray-900 mb-3">Joint Business Planning</h3>
            <p class="text-gray-600 text-sm leading-relaxed">Aligning with retailer growth trajectories and profitability goals to ensure "Carefully Deliberated" approvals.</p>
          </div>
        </div>
      </section>

      <!-- Analytics Section -->
      <section class="py-16 text-center bg-gray-50 rounded-2xl">
         <h2 class="text-2xl font-bold text-gray-900 mb-6">Turn "Ghost Inventory" into Data</h2>
         <p class="text-gray-600 mb-8 max-w-2xl mx-auto">Every interaction logged. Every missed volume accounted for. Transform your Weekly Review.</p>
         <button class="text-blue-600 font-medium hover:text-blue-800 text-sm">
           View Analytics Dashboard →
         </button>
      </section>

    </main>

    <!-- Footer -->
    <footer class="bg-white border-t border-gray-100 py-12">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 flex flex-col md:flex-row justify-between items-center">
        <div class="mb-4 md:mb-0 flex items-center">
          <img src="/amore-labs-logo.png" alt="Amore Labs" class="h-8 w-auto mr-2" />
          <span class="font-bold text-gray-900">Amore Labs</span>
        </div>
        <div class="flex space-x-8 text-sm text-gray-500">
          <a href="#" class="hover:text-gray-900">Joint Business Planning</a>
          <a href="#" class="hover:text-gray-900">Privacy</a>
          <a href="#" class="hover:text-gray-900">Contact D&S Planning</a>
        </div>
        <div class="mt-4 md:mt-0 text-sm text-gray-400">© 2026 Amore Labs Inc.</div>
      </div>
    </footer>

    <!-- Booking Modal -->
    <BookingModal :is-open="showBooking" @close="showBooking = false" />
  </div>
</template>
