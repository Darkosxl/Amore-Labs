<script setup lang="ts">
const emit = defineEmits<{
  (e: 'openBooking'): void
}>()

// Focused list of Enterprise/Productivity integrations based on the user's request
const integrations = [
  { name: "SAP", url: "https://upload.wikimedia.org/wikipedia/commons/5/59/SAP_2011_logo.svg" },
  { name: "Oracle", url: "https://upload.wikimedia.org/wikipedia/commons/5/50/Oracle_logo.svg" },
  { name: "Microsoft Teams", url: "https://upload.wikimedia.org/wikipedia/commons/thumb/0/07/Microsoft_Office_Teams_%282025%E2%80%93present%29.svg/250px-Microsoft_Office_Teams_%282025%E2%80%93present%29.svg.png" },
  { name: "Slack", url: "https://www.vectorlogo.zone/logos/slack/slack-icon.svg" },
  { name: "Notion", url: "https://upload.wikimedia.org/wikipedia/commons/4/45/Notion_app_logo.png" },
  { name: "Google Drive", url: "https://upload.wikimedia.org/wikipedia/commons/1/12/Google_Drive_icon_%282020%29.svg" },
  { name: "ServiceNow", url: "https://www.vectorlogo.zone/logos/servicenow/servicenow-icon.svg" },
  { name: "Salesforce", url: "https://www.vectorlogo.zone/logos/salesforce/salesforce-icon.svg" },
  { name: "HubSpot", url: "https://www.vectorlogo.zone/logos/hubspot/hubspot-icon.svg" },
  { name: "Coupa", url: "https://yt3.googleusercontent.com/OqEm0R9FC3KDct3HTTzK8TpQpYDQkCkPzh3AKydL6ptkxeZHS076-HcaAyAXqaD80BKQy2HlNmU=s900-c-k-c0x00ffffff-no-rj" },
  { name: "Workday", url: "https://www.vectorlogo.zone/logos/workday/workday-icon.svg" }
]

const octagonalClip = "polygon(30% 0%, 70% 0%, 100% 30%, 100% 70%, 70% 100%, 30% 100%, 0% 70%, 0% 30%)"
</script>

<template>
  <section class="max-w-7xl mx-auto my-20 px-6 grid md:grid-cols-2 gap-16 items-center border border-slate-800 bg-slate-900 p-10 rounded-3xl relative overflow-hidden">
    <!-- Background Glow -->
    <div class="absolute top-0 right-0 -mr-20 -mt-20 w-96 h-96 bg-blue-900/10 rounded-full blur-3xl pointer-events-none"></div>

    <!-- Left Side -->
    <div class="relative z-10">
      <p class="uppercase text-xs font-bold tracking-widest text-blue-400 mb-3">
        Ecosystem
      </p>
      <h2 class="text-5xl font-bold mb-6 text-white leading-tight">
        Unify Your<br/> Supply Chain
      </h2>
      <p class="text-slate-400 mb-8 text-lg leading-relaxed">
        Connect to SAP, Oracle, NetSuite, and your entire stack.
        Sync data seamlessly with enterprise-grade security (SOC II Type 1).
        Zero IT complexity required.
      </p>
      <div class="flex gap-4">
        <button @click="emit('openBooking')" class="bg-white text-slate-900 px-8 py-3.5 rounded-lg font-bold hover:bg-slate-200 transition-colors shadow-[0_0_20px_rgba(255,255,255,0.1)]">
          Book a Demo
        </button>
      </div>
    </div>

    <!-- Right Side -->
    <div class="grid grid-cols-4 gap-4 relative z-10">
      <div
        v-for="(item, idx) in integrations"
        :key="idx"
        class="relative w-20 h-20 bg-slate-800 shadow-lg border border-slate-700/50 hover:border-blue-500/50 transition-all duration-300 hover:-translate-y-1 hover:shadow-blue-500/10 group"
        :style="{ clipPath: octagonalClip }"
      >
        <div class="w-full h-full flex items-center justify-center bg-slate-800/50 backdrop-blur-sm p-4">
            <img
              :src="item.url"
              :alt="item.name"
              class="object-contain w-full h-full transition-all duration-300"
            />
        </div>
      </div>
    </div>
  </section>
</template>
