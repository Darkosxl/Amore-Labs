<script setup lang="ts">
import Card from './ui/BaseCard.vue'

const conversation = [
  {
    role: 'agent',
    text: "Hey, we ran into issues in producing 500 units of X, we have 300 units of X and I can give you 200 units of Y, is that ok?",
    time: '13:38'
  },
  {
    role: 'user',
    text: "No, I need 500 units of X, I already have too much Y,",
    time: '13:42'
  },
  {
    role: 'agent',
    text: "I have 300 units of X and 100 units of X family size, if you accept also 100 units of Z I can offer a 5% discount for the next week's batch of X.",
    time: '13:45'
  },
  {
    role: 'user',
    text: "Deal 🤝",
    time: '13:50'
  }
]
</script>

<template>
  <section class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
    <!-- Main Dark Container -->
    <div class="bg-slate-900 rounded-3xl p-8 md:p-12 overflow-hidden relative shadow-2xl">
      <!-- Background Effects (Gradient/Glow) - Monochrome -->
      <div class="absolute top-0 right-0 -mr-20 -mt-20 w-96 h-96 bg-white/5 rounded-full blur-3xl pointer-events-none"></div>
      <div class="absolute bottom-0 left-0 -ml-20 -mb-20 w-96 h-96 bg-gray-500/10 rounded-full blur-3xl pointer-events-none"></div>

      <!-- Header Section -->
      <div class="flex flex-col md:flex-row justify-between items-center mb-12 relative z-10">
        <h2 class="text-2xl md:text-3xl font-bold text-white mb-6 md:mb-0">
          Seamless ERP Integration
        </h2>

        <!-- Partner Logos (Text/Placeholder for ERPs) -->
        <!-- Partner Logos (Removed) -->
        <div class="flex items-center space-x-6"></div>
      </div>

      <!-- Cards Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6 relative z-10">

        <!-- Card 1: Live Transcript -->
        <Card variant="glass" class="flex flex-col overflow-hidden h-96">
          <div class="px-5 py-4 border-b border-white/5 flex justify-between items-center bg-slate-800/80">
            <div class="font-medium text-sm text-slate-200">Live Transcript</div>
            <div class="flex items-center space-x-2">
              <span class="relative flex h-2 w-2">
                <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-green-400 opacity-75"></span>
                <span class="relative inline-flex rounded-full h-2 w-2 bg-green-500"></span>
              </span>
              <span class="text-xs text-green-400 font-mono">Call in Progress</span>
              <span class="text-xs text-slate-500 font-mono">01:38</span>
            </div>
          </div>

          <div class="flex-1 overflow-y-auto p-4 space-y-4 font-sans text-sm scrollbar-thin scrollbar-thumb-slate-700 scrollbar-track-transparent">
             <div v-for="(msg, i) in conversation" :key="i"
                  :class="['flex gap-3', msg.role === 'user' ? 'flex-row-reverse' : '']">

                <!-- Avatar -->
                <div :class="['w-8 h-8 rounded-full flex items-center justify-center flex-shrink-0 text-xs font-bold', msg.role === 'agent' ? 'bg-blue-600 text-white' : 'bg-slate-600 text-slate-200']">
                  {{ msg.role === 'agent' ? 'AI' : 'U' }}
                </div>

                <!-- Bubble -->
                <div :class="['p-3 rounded-2xl max-w-[80%]',
                   msg.role === 'agent' ? 'bg-slate-700/80 rounded-tl-sm text-slate-200' : 'bg-blue-600 rounded-tr-sm text-white']">
                   {{ msg.text }}
                </div>
             </div>
          </div>

          <div class="px-4 py-3 bg-slate-900/50 text-xs text-slate-500 border-t border-white/5 flex justify-between">
            <span>AI-powered transcription</span>
            <span>Real-time analysis</span>
          </div>
        </Card>

        <!-- Card 2: Calendar -->
        <Card variant="glass" class="p-6 flex flex-col h-96">
           <div class="flex justify-between items-center mb-6">
             <div class="font-bold text-white">March 2026</div>
             <div class="flex items-center space-x-2 text-xs">
                <span class="w-2 h-2 rounded-full bg-green-500"></span>
                <span class="text-green-400">Booking Confirmed</span>
             </div>
           </div>

           <!-- Calendar Grid Mockup -->
           <div class="w-full">
              <div class="grid grid-cols-7 text-xs text-slate-500 mb-2 text-center">
                <span>S</span><span>M</span><span>T</span><span>W</span><span>T</span><span>F</span><span>S</span>
              </div>
              <div class="grid grid-cols-7 gap-1 text-center text-sm font-medium text-slate-300">
                 <!-- Padding days -->
                 <span class="p-2 text-slate-600">25</span><span class="p-2 text-slate-600">26</span><span class="p-2 text-slate-600">27</span><span class="p-2 text-slate-600">28</span><span class="p-2 text-slate-600">29</span>
                 <span class="p-2">1</span><span class="p-2">2</span>
                 <span class="p-2">3</span><span class="p-2">4</span><span class="p-2">5</span><span class="p-2">6</span><span class="p-2">7</span><span class="p-2">8</span><span class="p-2">9</span>
                 <span class="p-2">10</span><span class="p-2">11</span><span class="p-2">12</span>
                 <!-- Highlighted Day -->
                 <span class="p-2 bg-blue-600 text-white rounded-lg shadow-md hover:bg-blue-500 transition-colors cursor-pointer">13</span>
                 <span class="p-2 bg-green-600/20 text-green-400 rounded-lg">14</span><span class="p-2">15</span><span class="p-2">16</span>
              </div>
           </div>

           <!-- Appointment Details -->
           <div class="mt-auto bg-slate-900/50 p-4 rounded-lg border border-slate-700/50">
             <div class="text-xs text-slate-400 mb-1">Upcoming Modification</div>
             <div class="flex justify-between items-center">
               <div class="font-bold text-white">Batched Order #549</div>
               <div class="text-xs bg-green-500/10 text-green-500 px-2 py-1 rounded">Synced</div>
             </div>
           </div>
        </Card>

        <!-- Card 3: Analytics -->
        <Card variant="glass" class="p-6 flex flex-col h-96">
           <div class="flex justify-between items-center mb-6">
             <div class="font-bold text-white">Performance Analytics</div>
             <div class="flex items-center space-x-1 text-xs text-green-400">
               <span class="w-2 h-2 rounded-full bg-green-500 animate-pulse"></span>
               <span>Live</span>
             </div>
           </div>

           <div class="grid grid-cols-2 gap-4 mb-8">
             <div class="bg-slate-700/50 p-4 rounded-xl border border-slate-600/30">
               <div class="text-xs text-blue-400 font-medium mb-1">Orders Fulfilled</div>
               <div class="text-2xl font-bold text-white">1,247</div>
               <div class="text-[10px] text-green-400 mt-1">+23% this month</div>
             </div>
             <div class="bg-slate-700/50 p-4 rounded-xl border border-slate-600/30">
               <div class="text-xs text-green-400 font-medium mb-1">On-Time In-Full</div>
               <div class="text-2xl font-bold text-white">98%</div>
               <div class="text-[10px] text-green-400 mt-1">up 3% from last week</div>
             </div>
           </div>

           <div>
             <div class="text-xs text-slate-400 mb-3">Booking Trends (Last 7 Days)</div>
             <div class="flex items-end space-x-2 h-24 pb-2">
                <div class="flex-1 bg-blue-500/20 hover:bg-blue-500 rounded-t-sm h-[40%] transition-all"></div>
                <div class="flex-1 bg-blue-500/20 hover:bg-blue-500 rounded-t-sm h-[60%] transition-all"></div>
                <div class="flex-1 bg-blue-600 rounded-t-sm h-[85%] shadow-lg shadow-blue-500/20"></div>
                <div class="flex-1 bg-blue-500/20 hover:bg-blue-500 rounded-t-sm h-[70%] transition-all"></div>
                <div class="flex-1 bg-blue-500/20 hover:bg-blue-500 rounded-t-sm h-[50%] transition-all"></div>
                <div class="flex-1 bg-blue-500/20 hover:bg-blue-500 rounded-t-sm h-[45%] transition-all"></div>
                <div class="flex-1 bg-blue-500/20 hover:bg-blue-500 rounded-t-sm h-[75%] transition-all"></div>
             </div>
             <div class="flex justify-between text-[10px] text-slate-500 uppercase font-mono mt-1">
               <span>M</span><span>T</span><span>W</span><span>T</span><span>F</span><span>S</span><span>S</span>
             </div>
           </div>
        </Card>

      </div>

      <!-- Footer Info -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mt-16 text-slate-300 relative z-10 border-t border-slate-700/50 pt-10">
        <div>
          <h3 class="text-white font-bold mb-2">Calls retailers in 10 seconds 24/7</h3>
          <p class="text-sm text-slate-400 leading-relaxed">
            Handle inbound supply checks, outbound negotiation campaigns, and scheduled follow-ups with instant response times around the clock.
          </p>
        </div>
        <div>
          <h3 class="text-white font-bold mb-2">Integrates into every major ERP</h3>
          <p class="text-sm text-slate-400 leading-relaxed">
            Seamlessly integrates with all major enterprise resource planning systems globally. Adjustments sync directly into your existing workflow.
          </p>
        </div>
        <div>
          <h3 class="text-white font-bold mb-2">Negotiates with Tribal Knowledge</h3>
          <p class="text-sm text-slate-400 leading-relaxed">
             Real results from supply chain partners using our Adversarial Agents. Converting stockouts into confirmed substitutions with substantial revenue protection.
          </p>
        </div>
      </div>

      <!-- CTA Link -->
      <div class="mt-12 relative z-10">
        <a href="#" class="inline-flex items-center text-sm font-medium text-blue-400 hover:text-blue-300 transition-colors group">
          See in Action - Live Demo
          <svg class="w-4 h-4 ml-1 transform group-hover:translate-x-1 transition-transform" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3" />
          </svg>
        </a>
      </div>

    </div>
  </section>
</template>

<style scoped>
/* Custom scrollbar for the chat */
.scrollbar-thin::-webkit-scrollbar {
  width: 4px;
}
.scrollbar-thumb-slate-700::-webkit-scrollbar-thumb {
  background-color: #334155;
  border-radius: 4px;
}
</style>
