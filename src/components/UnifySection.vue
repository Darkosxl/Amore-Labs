<script setup lang="ts">
// Placeholder logic for Orbit animation if needed, mostly doing via CSS
</script>

<template>
  <section class="py-24 bg-slate-900 border-t border-slate-800">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-16 items-center">

        <!-- Left Column: Copy -->
        <div class="space-y-8">
          <h2 class="text-4xl md:text-5xl font-bold text-white tracking-tight leading-tight">
            Unify Your<br />
            Supply Chain<br />
            Systems and<br />
            Data. Safely.
          </h2>

          <div class="space-y-6">
            <!-- Feature 1 -->
            <div class="flex items-start">
              <div class="flex-shrink-0 mt-1">
                <div class="flex items-center justify-center h-6 w-6 rounded-full bg-blue-500/20 text-blue-400">
                  <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7" />
                  </svg>
                </div>
              </div>
              <div class="ml-4">
                <h3 class="text-lg font-bold text-white">Seamless Connections</h3>
                <p class="mt-2 text-slate-400 text-sm leading-relaxed">
                  Connect to SAP, Oracle, NetSuite, Teams, Slack, and more in minutes.
                  Sync data with PowerBI and ERPs instantly. Quick setup with zero IT complexity.
                </p>
              </div>
            </div>

            <!-- Feature 2 -->
            <div class="flex items-start">
              <div class="flex-shrink-0 mt-1">
                <div class="flex items-center justify-center h-6 w-6 rounded-full bg-blue-500/20 text-blue-400">
                  <svg class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7" />
                  </svg>
                </div>
              </div>
              <div class="ml-4">
                <h3 class="text-lg font-bold text-white">Enterprise-grade security</h3>
                <p class="mt-2 text-slate-400 text-sm leading-relaxed">
                  SOC II Type 1 certified. End-to-end encryption exceeding industry standards.
                  Your data remains yours, always protected.
                </p>
              </div>
            </div>
          </div>

          <div class="pt-4">
            <button class="bg-white text-slate-900 px-8 py-3 rounded-md text-sm font-bold hover:bg-slate-200 transition-colors shadow-lg shadow-white/10">
              Book a Demo
            </button>
          </div>
        </div>

        <!-- Right Column: Orbit Visual -->
        <div class="relative flex items-center justify-center min-h-[500px]">
           <!-- Background Glow -->
           <div class="absolute w-[500px] h-[500px] bg-blue-500/5 rounded-full blur-3xl pointer-events-none"></div>

           <!-- Central Badge -->
           <div class="relative z-20 w-32 h-32 bg-gradient-to-br from-blue-600 to-blue-800 rounded-full flex items-center justify-center shadow-2xl shadow-blue-500/20 border-4 border-slate-900">
             <div class="text-center text-white">
                <div class="font-bold text-xl leading-none">AICPA</div>
                <div class="text-xs font-medium tracking-widest opacity-80 mt-1">SOC</div>
                <div class="w-8 h-1 bg-white/20 mx-auto mt-2 rounded-full"></div>
             </div>
           </div>

           <!-- Orbit Ring 1 (Small) -->
           <div class="absolute z-10 w-64 h-64 border border-slate-800 rounded-full animate-spin-slow">
              <!-- Item 1 (Top) -->
              <div class="absolute top-0 left-1/2 -translate-x-1/2 -translate-y-1/2 w-12 h-12 bg-slate-800 border border-slate-700 rounded-xl flex items-center justify-center shadow-lg transform -rotate-0">
                  <span class="text-[10px] font-bold text-slate-300">Slack</span>
              </div>
              <!-- Item 2 (Bottom Right) -->
              <div class="absolute bottom-4 right-4 w-12 h-12 bg-slate-800 border border-slate-700 rounded-xl flex items-center justify-center shadow-lg">
                  <span class="text-[10px] font-bold text-slate-300">Teams</span>
              </div>
               <!-- Item 3 (Bottom Left) -->
              <div class="absolute bottom-4 left-4 w-12 h-12 bg-slate-800 border border-slate-700 rounded-xl flex items-center justify-center shadow-lg">
                  <span class="text-[10px] font-bold text-slate-300">Drive</span>
              </div>
           </div>

           <!-- Orbit Ring 2 (Large) -->
           <div class="absolute z-10 w-96 h-96 border border-slate-700/50 rounded-full animate-spin-reverse-slow">
              <!-- Item 1 (Right) -->
              <div class="absolute top-1/2 right-0 translate-x-1/2 -translate-y-1/2 w-16 h-16 bg-white border border-slate-200 rounded-2xl flex items-center justify-center shadow-xl">
                  <span class="font-bold text-xs text-blue-900 tracking-tighter">SAP</span>
              </div>
              <!-- Item 2 (Left) -->
              <div class="absolute top-1/2 left-0 -translate-x-1/2 -translate-y-1/2 w-16 h-16 bg-white border border-slate-200 rounded-2xl flex items-center justify-center shadow-xl">
                  <span class="font-bold text-xs text-red-600 tracking-tighter">ORACLE</span>
              </div>
              <!-- Item 3 (Top) -->
               <div class="absolute top-0 left-1/2 -translate-x-1/2 -translate-y-1/2 w-14 h-14 bg-slate-800 border border-slate-600 rounded-2xl flex items-center justify-center shadow-xl">
                  <span class="font-bold text-[10px] text-slate-300 italic">NetSuite</span>
              </div>
              <!-- Item 4 (Bottom) -->
               <div class="absolute bottom-0 left-1/2 -translate-x-1/2 translate-y-1/2 w-14 h-14 bg-slate-800 border border-slate-600 rounded-2xl flex items-center justify-center shadow-xl">
                  <span class="font-bold text-[10px] text-slate-300">Notion</span>
              </div>
           </div>

           <!-- Connecting Lines (Decorative) -->
           <svg class="absolute w-full h-full pointer-events-none opacity-20" viewBox="0 0 500 500">
              <path d="M250 250 L 400 250" stroke="white" stroke-dasharray="4 4" />
              <path d="M250 250 L 100 250" stroke="white" stroke-dasharray="4 4" />
           </svg>

        </div>

      </div>
    </div>
  </section>
</template>

<style scoped>
.animate-spin-slow {
  animation: spin 20s linear infinite;
}
.animate-spin-reverse-slow {
  animation: spin-reverse 25s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}
@keyframes spin-reverse {
  from { transform: rotate(360deg); }
  to { transform: rotate(0deg); }
}

/* Pause on hover for easier reading */
.group:hover .animate-spin-slow,
.group:hover .animate-spin-reverse-slow {
  animation-play-state: paused;
}
</style>
