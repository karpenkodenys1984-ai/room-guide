<script setup lang="ts">
import { ref, nextTick, inject } from "vue"
import { Handle, Position, useVueFlow, type NodeProps } from '@vue-flow/core'
import type { RoomData } from '../../types/graph';

const props = defineProps<NodeProps<RoomData>>()

const { updateNodeData } = useVueFlow()

const editing = ref(false)

const inputRef = ref<HTMLInputElement | null>(null)

const updateNodeLabel = inject<(id: string, label: string, x: number, y: number) => Promise<void>>('updateNodeLabel')

function startEdit() {
  editing.value = true
  nextTick(() => {
    inputRef.value?.focus()
    inputRef.value?.select()
  })
}

function stopEdit(e: Event) {
  const value = (e.target as HTMLInputElement).value.trim()
  if (value) {
    updateNodeData(props.id, { label: value })
    updateNodeLabel?.(props.id, value, props.position.x, props.position.y)
  }
  editing.value = false
}

function onKeydown(e: KeyboardEvent) {
  if (e.key === 'Enter') (e.target as HTMLInputElement).blur()
  if (e.key === 'Escape') editing.value = false
}

</script>

<template>
  <div class="room-node" @dblclick="startEdit">
    <Handle type="target" :position="Position.Left" />
    <input
      v-if="editing"
      ref="inputRef"
      class="label-input"
      :value="data.label"
      @blur="stopEdit"
      @keydown="onKeydown"
    />
    <span v-else>{{ data.label }}</span>
    <Handle type="source" :position="Position.Right" />
  </div>
</template>

<style scoped>
.room-node {
  background: white;
  border: 1.5px solid #378add;
  border-radius: 8px;
  padding: 6px 14px;
  font-size: 13px;
  cursor: default;
  color: #1e1e2e;
}

.label-input {
  border: none;
  outline: none;
  font-size: 13px;
  color: #1e1e2e;
  background: transparent;
  width: 40%;
  min-width: 60px;
}
</style>