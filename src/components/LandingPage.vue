<script setup lang="ts">
const navItems = [
  'Live Action Workflow',
  'The Metrics',
  'Negotiation Engine',
  'Integrations'
]

const workflowSteps = [
  { step: '01', title: 'The Trigger', description: 'D&S reviews retailer In-Stock measures' },
  { step: '02', title: 'Referral', description: 'Route requirements to CSC' },
  { step: '03', title: 'The Pitch', description: 'AI proposes offers to maximize fulfillment' },
  { step: '04', title: 'The Counter', description: 'Retailer responses reviewed by CSC' },
  { step: '05', title: 'Execution', description: 'Automated RFX & PO generation' },
  { step: '06', title: 'Visibility', description: 'Real-time sync of Reason Codes' }
]

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
        <div class="flex justify-between items-center h-20">
          <div class="flex-shrink-0 flex items-center">
             <!-- Using text logo to match previous clean style, or existing image if present -->
             <span class="font-bold text-xl tracking-tight">OTIF<span class="text-blue-600">.AI</span></span>
          </div>

          <nav class="hidden md:flex space-x-8">
            <a
              v-for="item in navItems"
              :key="item"
              href="#"
              class="text-[11px] font-medium text-gray-500 hover:text-gray-900 tracking-wider uppercase"
            >
              {{ item }}
            </a>
          </nav>

          <div class="flex items-center space-x-4">
            <a href="#" class="text-sm font-medium text-gray-500 hover:text-gray-900">
              Hear a Demo
            </a>
            <button class="bg-black text-white px-5 py-2.5 rounded-md text-sm font-medium hover:bg-gray-800 transition-colors">
              Deploy Agents
            </button>
          </div>
        </div>
      </div>
    </header>

    <!-- Hero Section -->
    <main class="pt-32 pb-16 px-4 sm:px-6 lg:px-8 max-w-7xl mx-auto">
      <div class="text-center mb-24">
        <div class="inline-flex items-center space-x-2 bg-red-50 px-3 py-1 rounded-full mb-8 border border-red-100">
           <span class="text-red-600 text-xs font-bold tracking-wider uppercase">Avoid the 3% COGS Non-Compliance Fine</span>
        </div>

        <h1 class="text-4xl sm:text-5xl md:text-6xl font-bold tracking-tight text-gray-900 mb-6">
          Protect Your 98% OTIF Scorecard<br class="hidden sm:block" />
          with Adversarial Voice Agents
        </h1>

        <p class="text-lg sm:text-xl text-gray-600 max-w-3xl mx-auto mb-10">
          Don't just accept the fine. When production lines break, our AI negotiates substitutions, prevents on-shelf gaps, and secures Best-in-Class status before the truck leaves.
        </p>

        <div class="flex justify-center">
          <button class="bg-gray-900 text-white px-8 py-3 rounded-md text-base font-medium hover:bg-gray-800 transition-colors">
            Deploy Agents
          </button>
        </div>
      </div>

      <!-- Live Action Workflow -->
      <section class="py-20 border-t border-gray-100">
        <h2 class="text-center text-xs font-bold text-gray-400 uppercase tracking-widest mb-16">Live Action Workflow</h2>

        <div class="grid grid-cols-1 md:grid-cols-6 gap-8 text-center">
          <div v-for="(item, index) in workflowSteps" :key="index" class="relative">
             <div class="text-xs font-mono text-gray-400 mb-2">{{ item.step }}</div>
             <div class="w-12 h-12 mx-auto bg-gray-50 rounded-full flex items-center justify-center mb-4 border border-gray-100">
               <span class="text-xl">↓</span>
             </div>
             <h3 class="text-sm font-bold text-gray-900 mb-1">{{ item.title }}</h3>
             <p class="text-xs text-gray-500 px-2">{{ item.description }}</p>
          </div>
        </div>
      </section>

      <!-- Three Views of Failure -->
      <section class="py-20 border-t border-gray-100">
        <div class="text-center mb-16">
          <h2 class="text-3xl font-bold text-gray-900 mb-4">Three Views of Failure</h2>
          <p class="text-gray-600">Why the AI is a strategic necessity.</p>
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
                 <div class="w-8 h-8 rounded-full bg-blue-100 flex items-center justify-center text-blue-600 text-xs font-bold">AI</div>
                 <div class="bg-white p-3 rounded-tr-xl rounded-br-xl rounded-bl-xl shadow-sm text-sm text-gray-700">
                   Connecting to Retailer Buyer...
                 </div>
               </div>

               <div class="flex items-start space-x-3">
                 <div class="w-8 h-8 rounded-full bg-blue-100 flex items-center justify-center text-blue-600 text-xs font-bold">AI</div>
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
        <div class="mb-4 md:mb-0 font-bold text-gray-900">OTIF<span class="text-blue-600">.AI</span></div>
        <div class="flex space-x-8 text-sm text-gray-500">
          <a href="#" class="hover:text-gray-900">Joint Business Planning</a>
          <a href="#" class="hover:text-gray-900">Privacy</a>
          <a href="#" class="hover:text-gray-900">Contact D&S Planning</a>
        </div>
        <div class="mt-4 md:mt-0 text-sm text-gray-400">© 2026 OTIF AI Agents.</div>
      </div>
    </footer>
  </div>
</template>
