<script setup lang="ts">
import type { FlowNode, RoomData } from '../../types/graph.js'
import { computed } from 'vue';

const props = defineProps<{
  nodes: FlowNode[]
}>()

const emit = defineEmits<{
  delete: [nodeId: string]
}>()

function isRoomData(data: FlowNode['data']): data is RoomData {
  return 'label' in data
}

const deletableNodes = computed(() =>
  props.nodes.filter((n): n is FlowNode & { data: RoomData } => isRoomData(n.data))
)


const selectedNodeId = defineModel<string | null>('selectedNodeId', { default: null })

function handleDelete() {
  if (!selectedNodeId.value) return
  emit('delete', selectedNodeId.value)
}
</script>

<template>
  <div class="delete-panel">
    <select v-model="selectedNodeId">
      <option disabled value="">Choose Node</option>
      <option v-for="n in deletableNodes" :key="n.id" :value="n.id">
        {{ n.data.label }} ({{ n.id }})
      </option>
    </select>
    <button :disabled="!selectedNodeId" @click="handleDelete">Delete</button>
  </div>
</template>

<style scoped>
.delete-panel {
  display: flex;
  gap: 10px;
  align-items: center;
  padding: 10px 15px;
  background-color: #313244; /* surface0 */
  border: 1px solid #45475a; /* surface1 */
  border-radius: 8px;
  margin: 0 15px;
}

.delete-panel select {
  flex: 1;
  padding: 6px 10px;
  background-color: #1e1e2e; /* base */
  color: #cdd6f4; /* text */
  border: 1px solid #45475a;
  border-radius: 6px;
  font-size: 14px;
  outline: none;
  cursor: pointer;
  transition: border-color 0.15s ease;
}

.delete-panel select:hover {
  border-color: #6366f1;
}

.delete-panel select:focus {
  border-color: #6366f1;
  box-shadow: 0 0 0 2px rgba(99, 102, 241, 0.25);
}

.delete-panel select option {
  background-color: #1e1e2e;
  color: #cdd6f4;
}

.delete-panel button {
  padding: 6px 16px;
  background-color: #f38ba8; /* red */
  color: #1e1e2e;
  border: none;
  border-radius: 6px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.15s ease, opacity 0.15s ease;
  white-space: nowrap;
}

.delete-panel button:hover:not(:disabled) {
  background-color: #eb6f92;
}

.delete-panel button:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}
</style>