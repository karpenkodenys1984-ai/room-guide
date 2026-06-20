<script setup lang="ts">
import type { NodeProps } from '@vue-flow/core'

const props = defineProps<NodeProps<{ src: string; width: number; height: number }>>()
</script>

<template>
  <div :style="{ width: props.data.width + 'px', height: props.data.height + 'px' }">
    <img 
      :src="props.data.src" 
      :width="props.data.width" 
      :height="props.data.height" 
      style="display:block; opacity:0.7; pointer-events:none" 
    />
  </div>
</template>