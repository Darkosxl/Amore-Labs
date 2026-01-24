<script setup lang="ts">
const navItems = [
  { label: 'Why Adversarial Agents', targetId: 'seamless-erp-integration' },
  { label: 'How It Works', targetId: 'how-amore-labs-transforms' },
  { label: 'Use Cases', targetId: 'powering-manufacturing-supply-chains' },
  { label: 'Integrations', targetId: 'unify-your-supply-chain' }
]

import LiquidNavbar from './ui/LiquidNavbar.vue'
import IntegrationSection from './IntegrationSection.vue'
import IntegrationsGrid from './IntegrationsGrid.vue'
import BookingModal from './BookingModal.vue'
import ManufacturingPowerSection from './ManufacturingPowerSection.vue' // New Section
import GlassIcon from './ui/GlassIcon.vue'
import '@/assets/GlassIcons.css'
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
    <!-- Liquid Navbar -->
    <LiquidNavbar :nav-items="navItems" @open-booking="showBooking = true" />

    <!-- Hero Section -->
    <main id="hero" class="pt-32 pb-16 px-4 sm:px-6 lg:px-8 max-w-7xl mx-auto">
      <div class="text-center mb-24">

        <h1 class="text-4xl sm:text-5xl md:text-6xl font-bold tracking-tighter text-gray-900 mb-6">
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
      <IntegrationSection id="seamless-erp-integration" />

      <!-- New Transformation Section -->
      <section id="how-amore-labs-transforms" class="py-24 bg-white border-t border-gray-100">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">

           <!-- Relative-Unit Outlined Card Structure -->
           <div class="max-w-5xl mx-auto relative pt-[2vh]">

              <!-- Headline Content (Moved Outside) -->
              <div class="text-center pb-[5vh] text-gray-900 relative z-10">
                  <h3 class="text-[2.25rem] font-bold mb-[1.5vh] tracking-tight">Enterprise-grade automation</h3>
                  <p class="text-gray-600 max-w-2xl mx-auto text-[1.125rem] leading-relaxed">
                    From trigger to execution, our adversarial agents handle the complexity.
                    No training required. Just results.
                  </p>
              </div>

              <!-- Light Card Frame (Outer) -->
              <div class="relative w-full rounded-[3rem] p-[2%] shadow-2xl border border-gray-200 bg-gradient-to-br from-white via-white to-slate-50">

                   <!-- Dark Card content (Inner Overlay) -->
                   <div class="rounded-[2rem] shadow-xl p-[4vw] md:p-[6%] relative z-10 overflow-hidden border border-slate-800"
                        style="background: radial-gradient(circle at 0% 0%, #000000, #505050, #000000), linear-gradient(180deg, #a6a1a1, #000000); background-blend-mode: overlay, normal;">
                       <!-- Background Grid -->
                       <div class="absolute inset-0 opacity-20 pointer-events-none"
                            style="background-image: radial-gradient(#ffffff 1px, transparent 1px); background-size: 32px 32px;">
                       </div>
                       <div class="grid grid-cols-1 md:grid-cols-3 gap-y-[10vh] gap-x-[5%] relative max-w-full mx-auto">

                           <!-- Step 1 -->
                           <div class="flex flex-col items-center text-center relative group">
                             <GlassIcon color="blue" customClass="mb-[2vh] transition-transform group-hover:scale-110">
                               <svg class="w-[50%] h-[50%] text-blue-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                               </svg>
                             </GlassIcon>
                             <div class="text-[0.65rem] font-bold text-gray-400 uppercase tracking-widest mb-[0.5vh]">Step 1</div>
                             <h4 class="text-[1.25rem] font-bold text-white mb-[1vh]">The Trigger</h4>
                             <p class="text-[0.875rem] text-slate-300 leading-relaxed max-w-[15rem]">D&S reviews retailer In-Stock measures and detects stockouts.</p>
                           </div>

                           <!-- Step 2 -->
                           <div class="flex flex-col items-center text-center relative group">
                             <GlassIcon color="purple" customClass="mb-[2vh] transition-transform group-hover:scale-110">
                               <svg class="w-[50%] h-[50%] text-purple-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
                               </svg>
                             </GlassIcon>
                             <div class="text-[0.65rem] font-bold text-gray-400 uppercase tracking-widest mb-[0.5vh]">Step 2</div>
                             <h4 class="text-[1.25rem] font-bold text-white mb-[1vh]">Referral</h4>
                             <p class="text-[0.875rem] text-slate-300 leading-relaxed max-w-[15rem]">Route requirements to CSC Adversarial Agents.</p>
                           </div>

                           <!-- Step 3 -->
                           <div class="flex flex-col items-center text-center relative group">
                             <GlassIcon color="orange" customClass="mb-[2vh] transition-transform group-hover:scale-110">
                               <svg class="w-[50%] h-[50%] text-orange-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7h12m0 0l-4-4m4 4l-4 4m0 6H4m0 0l4 4m-4-4l4-4" />
                               </svg>
                             </GlassIcon>
                             <div class="text-[0.65rem] font-bold text-gray-400 uppercase tracking-widest mb-[0.5vh]">Step 3</div>
                             <h4 class="text-[1.25rem] font-bold text-white mb-[1vh]">The Pitch</h4>
                             <p class="text-[0.875rem] text-slate-300 leading-relaxed max-w-[15rem]">Adversarial Agent proposes offers to maximize fulfillment.</p>
                           </div>

                           <!-- Step 6 (Bottom Left) -->
                           <div class="flex flex-col items-center text-center relative group order-last md:order-none">
                             <GlassIcon color="indigo" customClass="mb-[2vh] transition-transform group-hover:scale-110">
                                <svg class="w-[50%] h-[50%] text-indigo-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 002 2h2a2 2 0 002-2z" />
                                </svg>
                             </GlassIcon>
                             <div class="text-[0.65rem] font-bold text-gray-400 uppercase tracking-widest mb-[0.5vh]">Step 6</div>
                             <h4 class="text-[1.25rem] font-bold text-white mb-[1vh]">Visibility</h4>
                             <p class="text-[0.875rem] text-slate-300 leading-relaxed max-w-[15rem]">Real-time sync of Reason Codes and delivery status.</p>
                           </div>

                           <!-- Step 5 (Bottom Middle) -->
                           <div class="flex flex-col items-center text-center relative group order-5 md:order-none">
                             <GlassIcon color="green" customClass="mb-[2vh] transition-transform group-hover:scale-110">
                               <svg class="w-[50%] h-[50%] text-green-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                 <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                               </svg>
                             </GlassIcon>
                             <div class="text-[0.65rem] font-bold text-gray-400 uppercase tracking-widest mb-[0.5vh]">Step 5</div>
                             <h4 class="text-[1.25rem] font-bold text-white mb-[1vh]">Execution</h4>
                             <p class="text-[0.875rem] text-slate-300 leading-relaxed max-w-[15rem]">Automated RFX & PO generation.</p>
                           </div>

                           <!-- Step 4 (Bottom Right) -->
                           <div class="flex flex-col items-center text-center relative group order-4 md:order-none">
                             <GlassIcon color="yellow" customClass="mb-[2vh] transition-transform group-hover:scale-110">
                                <svg class="w-[50%] h-[50%] text-yellow-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                                </svg>
                             </GlassIcon>
                             <div class="text-[0.65rem] font-bold text-gray-400 uppercase tracking-widest mb-[0.5vh]">Step 4</div>
                             <h4 class="text-[1.25rem] font-bold text-white mb-[1vh]">The Counter</h4>
                             <p class="text-[0.875rem] text-slate-300 leading-relaxed max-w-[15rem]">Retailer responses reviewed by CSC.</p>
                           </div>

                       </div>
                   </div>

              </div>
           </div>

        </div>
      </section>

      <!-- Manufacturing Power Section (New - moved below) -->
      <ManufacturingPowerSection id="powering-manufacturing-supply-chains" />

      <!-- Integrations Grid (Octagonal Logos) -->
      <IntegrationsGrid id="unify-your-supply-chain" @openBooking="showBooking = true" />
    </main>

    <!-- Footer -->
    <footer class="bg-white border-t border-gray-100 py-12">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 flex flex-col md:flex-row justify-between items-center">
        <div class="mb-4 md:mb-0 flex items-center">
          <img src="/amore-labs-logo.png" alt="Amore Labs" class="h-8 w-auto mr-2" />
          <span class="font-bold text-gray-900">Amore Labs</span>
        </div>
        <div class="mt-4 md:mt-0 text-sm text-gray-400">© 2026 Amore Labs Inc.</div>
      </div>
    </footer>

    <!-- Booking Modal -->
    <BookingModal :is-open="showBooking" @close="showBooking = false" />
  </div>
</template>
