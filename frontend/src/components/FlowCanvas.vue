<script setup lang="ts">
import { VueFlow, useVueFlow, type NodeTypesObject, type NodeDragEvent, type Node as VueFlowNode } from '@vue-flow/core'
import '@vue-flow/core/dist/style.css'
import RoomNode from './nodes/RoomNode.vue'
import { FlowNode, FlowEdge } from '../types/graph.js'

import { markRaw } from 'vue'
import BackgroundNode from './nodes/BackgroundNode.vue'

const { screenToFlowCoordinate } = useVueFlow()

const props = defineProps<{
  backgroundImage?: string | null,
  nodes?: FlowNode[],
  edges?: FlowEdge[]

}>()

const emit = defineEmits<{
    onConnected: [],
    onNodeDragStop: [node: VueFlowNode],
}>();


const nodeTypes: NodeTypesObject = {
  room: markRaw(RoomNode),
  background: markRaw(BackgroundNode),
}

function onConnected() {
  emit("onConnected")
}

function onNodeDragStop({ node }: NodeDragEvent) {
  if (node.type === 'background') return
  emit('onNodeDragStop', node)
}

</script>

<template>
  <div class="flow-wrapper">
      <VueFlow
        :nodes="nodes"
        :edges="edges"
        :node-types="nodeTypes"
        fit-view-on-init
        @connect="onConnected"
        @node-drag-stop="onNodeDragStop"
  >
  </VueFlow>
  </div>
</template>

<style scoped>
.flow-wrapper {
  width: 100%;
  height: 100%;
  background: white;
}
.floor-plan-wrap {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}
.floor-plan {
 display: block;
  width: 100%;
  height: 100%;
  object-fit: contain;
  opacity: 0.90;
}
</style>