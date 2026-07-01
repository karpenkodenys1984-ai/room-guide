<script setup lang="ts">
import type { FlowNode } from '../../types/graph.js'

const props = defineProps<{
  nodes: FlowNode[]
}>()

const emit = defineEmits<{
  delete: [nodeId: string]
}>()

const selectedNodeId = defineModel<string | null>('selectedNodeId', { default: null })

function handleDelete() {
  if (!selectedNodeId.value) return
  emit('delete', selectedNodeId.value)
}
</script>

<template>
  <div class="delete-panel">
    <select v-model="selectedNodeId">
      <option disabled value="">Выберите ноду</option>
      <option v-for="n in nodes" :key="n.id" :value="n.id">
        {{ n.data.label }} ({{ n.id }})
      </option>
    </select>
    <button :disabled="!selectedNodeId" @click="handleDelete">Удалить</button>
  </div>
</template>

<style scoped>
.delete-panel {
  display: flex;
  gap: 8px;
  align-items: center;
  padding: 8px 15px;
}
</style>