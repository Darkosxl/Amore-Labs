<script setup lang="ts">
import { computed } from 'vue';

interface Props {
  color?: string;
  label?: string;
  customClass?: string;
}

const props = withDefaults(defineProps<Props>(), {
  color: 'blue',
  label: '',
  customClass: '',
});

const gradientMapping: Record<string, string> = {
  blue: 'linear-gradient(hsl(223, 90%, 50%), hsl(208, 90%, 50%))',
  purple: 'linear-gradient(hsl(283, 90%, 50%), hsl(268, 90%, 50%))',
  red: 'linear-gradient(hsl(3, 90%, 50%), hsl(348, 90%, 50%))',
  indigo: 'linear-gradient(hsl(253, 90%, 50%), hsl(238, 90%, 50%))',
  orange: 'linear-gradient(hsl(43, 90%, 50%), hsl(28, 90%, 50%))',
  green: 'linear-gradient(hsl(123, 90%, 40%), hsl(108, 90%, 40%))',
  // Added yellow for compatibility with existing design
  yellow: 'linear-gradient(hsl(50, 90%, 50%), hsl(35, 90%, 50%))',
};

const backgroundStyle = computed(() => {
  const gradient = gradientMapping[props.color];
  if (gradient) {
    return { background: gradient };
  }
  return { background: props.color };
});
</script>

<template>
  <button :class="['icon-btn', customClass]" :aria-label="label" type="button">
    <!-- Back layer (Background/Gradient) -->
    <span class="icon-btn__back" :style="backgroundStyle"></span>

    <!-- Front layer (Glass + Icon) -->
    <span class="icon-btn__front">
      <span class="icon-btn__icon" aria-hidden="true">
        <slot></slot>
      </span>
    </span>

    <!-- Label (Optional) -->
    <span v-if="label" class="icon-btn__label">{{ label }}</span>
  </button>
</template>

<style scoped>
/* Scoped styles need to be careful with global classes from GlassIcons.css
   Since GlassIcons.css is imported globally or in parent, we rely on classes matching.
   However, for safety / self-containment, it's often better to import the CSS here or rely on global import.
   Assuming global import (or import in LandingPage) for now to match user instruction pattern.
*/
</style>
