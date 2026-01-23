<script setup lang="ts">
// Placeholder specific logic if needed
</script>

<template>
  <section class="py-24 bg-white relative overflow-hidden">
    <!-- Decorative Blob -->
    <div class="absolute top-0 left-0 w-[500px] h-[500px] bg-blue-50/50 rounded-full blur-3xl -translate-x-1/2 -translate-y-1/2 pointer-events-none"></div>

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-16 items-stretch">

        <!-- Left Side: Dark UI Mockup Visual -->
        <div class="relative group">
           <!-- Card Container -->
           <div class="relative rounded-3xl p-6 sm:p-10 shadow-2xl border border-slate-800 overflow-hidden h-full flex flex-col justify-center"
                style="background: radial-gradient(circle at 0% 0%, #000000, #505050, #000000), linear-gradient(180deg, #a6a1a1, #000000); background-blend-mode: overlay, normal;">

              <!-- Background Grid -->
              <div class="absolute inset-0 opacity-20"
                   style="background-image: radial-gradient(#ffffff 1px, transparent 1px); background-size: 24px 24px;">
              </div>

              <!-- Header -->
              <div class="flex items-center justify-between mb-8 relative z-10">
                 <h3 class="text-white font-bold text-lg">AI Resolution Workflow</h3>
                 <div class="flex items-center space-x-2">
                    <span class="w-2 h-2 bg-green-500 rounded-full animate-pulse"></span>
                    <span class="text-xs text-slate-400 font-mono">LIVE AUTOMATION</span>
                 </div>
              </div>

              <!-- Workflow Nodes -->
              <div class="space-y-6 relative z-10">

                  <!-- Node 1: Stockout Trigger -->
                  <div class="flex items-center space-x-4">
                      <div class="w-10 h-10 rounded-xl bg-red-500/20 flex items-center justify-center border border-red-500/30">
                          <svg class="w-5 h-5 text-red-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                          </svg>
                      </div>
                      <div class="bg-slate-800/80 backdrop-blur border border-slate-700 rounded-lg p-3 flex-1 shadow-lg">
                          <div class="text-xs text-slate-400 uppercase tracking-wider mb-1">Trigger Event</div>
                          <div class="text-white font-medium text-sm">Predicted Stockout: SKU-7782</div>
                          <div class="text-xs text-red-400 mt-1">Risk: 3% COGS Fine Immediate</div>
                      </div>
                  </div>

                  <!-- Connecting Line -->
                  <div class="h-8 w-0.5 bg-slate-700 ml-5"></div>

                  <!-- Node 2: Resolution Action -->
                  <div class="flex items-center space-x-4">
                      <div class="w-10 h-10 rounded-xl bg-blue-500/20 flex items-center justify-center border border-blue-500/30">
                          <svg class="w-5 h-5 text-blue-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                          </svg>
                      </div>
                      <div class="bg-slate-800/80 backdrop-blur border border-slate-700 rounded-lg p-3 flex-1 shadow-lg">
                           <div class="text-xs text-slate-400 uppercase tracking-wider mb-1">AI Action</div>
                           <div class="text-white font-medium text-sm">Drafted Substitution Proposal</div>
                           <div class="flex items-center mt-2 space-x-2">
                               <div class="h-1.5 flex-1 bg-slate-700 rounded-full overflow-hidden">
                                   <div class="h-full bg-blue-500 w-3/4"></div>
                               </div>
                               <span class="text-[10px] text-blue-400">92% Confidence</span>
                           </div>
                      </div>
                  </div>

                  <!-- Connecting Line -->
                  <div class="h-8 w-0.5 bg-slate-700 ml-5 opacity-50"></div>

                  <!-- Node 3: Pending -->
                  <div class="flex items-center space-x-4 opacity-70">
                      <div class="w-10 h-10 rounded-xl bg-slate-800 flex items-center justify-center border border-slate-700 border-dashed">
                          <div class="w-4 h-4 rounded-full border-2 border-slate-600 border-t-slate-400 animate-spin"></div>
                      </div>
                       <div class="text-sm text-slate-500">Waiting for retailer approval...</div>
                  </div>

              </div>
           </div>

           <!-- Decorative blurry glow behind -->
           <div class="absolute -inset-4 bg-blue-600/20 rounded-[2.5rem] blur-2xl -z-10 group-hover:bg-blue-600/30 transition-colors duration-1000"></div>
        </div>

        <!-- Right Side: Content -->
        <div>
          <p class="text-blue-600 font-bold uppercase tracking-widest text-xs mb-4">
             Manufacturing Specialized
          </p>
          <h2 class="text-4xl md:text-5xl font-bold text-gray-900 mb-6 leading-tight tracking-tighter">
            Powering<br/> Manufacturing<br/> Supply Chains
          </h2>
          <p class="text-lg text-gray-600 mb-10 leading-relaxed">
            We specialize in CPG and high-velocity inventory turnover.
            Generic solutions fail where speed matters. Automate the chaos of rapid fulfillment.
          </p>

          <div class="space-y-6">
             <!-- Use Case 1 -->
             <div class="flex items-start">
                 <div class="flex-shrink-0 mt-1">
                     <div class="w-8 h-8 rounded-full bg-red-100 flex items-center justify-center">
                         <svg class="w-4 h-4 text-red-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                           <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                         </svg>
                     </div>
                 </div>
                 <div class="ml-4">
                     <h4 class="text-gray-900 font-bold text-base">Avoid the 3% COGS Fine</h4>
                     <p class="text-gray-500 text-sm mt-1">Prevent non-compliance penalties before they happen with proactive gap detection.</p>
                 </div>
             </div>

             <!-- Use Case 2 -->
             <div class="flex items-start">
                 <div class="flex-shrink-0 mt-1">
                     <div class="w-8 h-8 rounded-full bg-blue-100 flex items-center justify-center">
                         <svg class="w-4 h-4 text-blue-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                           <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z" />
                         </svg>
                     </div>
                 </div>
                 <div class="ml-4">
                     <h4 class="text-gray-900 font-bold text-base">Rapid Inventory Turnover</h4>
                     <p class="text-gray-500 text-sm mt-1">Automated replenishment for high-velocity SKUs to keep production lines moving.</p>
                 </div>
             </div>

             <!-- Use Case 3 -->
             <div class="flex items-start">
                 <div class="flex-shrink-0 mt-1">
                     <div class="w-8 h-8 rounded-full bg-green-100 flex items-center justify-center">
                         <svg class="w-4 h-4 text-green-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                           <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                         </svg>
                     </div>
                 </div>
                 <div class="ml-4">
                     <h4 class="text-gray-900 font-bold text-base">Perfect Shelf Availability</h4>
                     <p class="text-gray-500 text-sm mt-1">Ensure product is on-shelf at the exact moment the consumer reaches for it.</p>
                 </div>
             </div>
          </div>

          <div class="mt-10">
             <button class="bg-gray-900 text-white px-8 py-3 rounded-md text-sm font-bold hover:bg-black transition-colors shadow-lg">
                See All Use Cases
             </button>
          </div>
        </div>

      </div>
    </div>
  </section>
</template>
