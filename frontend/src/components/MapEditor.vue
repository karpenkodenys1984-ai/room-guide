<script setup lang="ts">
import { SubTab, SubTabAction } from '../types/MapEditor/Tab';
import FlowCanvas from './FlowCanvas.vue';
import Tabs from './Map/Tabs.vue';
import NodeDeletePanel from './Map/NodeDeletePanel.vue';
import { onMounted, ref, provide, computed } from 'vue';
import { GetMapBackground, SaveMapBackground, SaveNode, GetAllNodes, UpdateNode, DeleteNode } from '../../wailsjs/go/main/App';
import { useLogger } from '../composables/useLogger';
import { XYPosition } from '@vue-flow/core';
import { useGraph } from '../composables/useGraph'
import type { Node as VueFlowNode } from '@vue-flow/core'
import { FlowNode, RoomData } from '../types/graph.js'


const { nodes, edges, onConnect } = useGraph()

const logger = useLogger()

const tabs  = [
    {
        name: "Map",
        icon: null,
        subTabs: [
            {
                name: "Upload",
                icon: null,
                action: SubTabAction.Upload
            },
        ]
    },
    {
        name: "Node",
        icon: null,
        subTabs: [
            {
                name: "Create",
                icon: null,
                action: SubTabAction.NodeCreate
            },
            {
                name: "Delete",
                icon: null,
                action: SubTabAction.NodeDelete
            }
        ]
    },
    {
        name: "Route",
        icon: null,
        subTabs: [
            {
                name: "Create",
                icon: null,
            },
            {
                name: "Delete",
                icon: null,
            }
        ]
    }
]

const loading = ref(false);

const actionHandlers: Record<SubTabAction, () => void> = {
    [SubTabAction.Upload]:      () => fileInput.value?.click(),
    [SubTabAction.NodeCreate]:  () => createNode(),
    [SubTabAction.NodeDelete]:  () => handleDeleteNode(),
    [SubTabAction.RouteCreate]: () => console.log('create route'),
    [SubTabAction.RouteDelete]: () => console.log('delete route'),
}

const fileInput = ref<HTMLInputElement | null>(null)

const backgroundImage = ref<string|null>(null)

const showDeletePanel = ref(false)
const selectedNodeId = ref<string | null>(null)

const deletableNodes = computed(() =>
  nodes.value.filter(n => n.type !== 'background')
)

function handleDeleteNode() {
  showDeletePanel.value = true
}

async function deleteSelectedNode(nodeId: string) {
  try {
    await DeleteNode(Number(nodeId))
    nodes.value = nodes.value.filter(n => n.id !== nodeId)
  } catch {
     logger.error('failed to delete node', { error: String(e) })
  }
}


function subTabChangeHandler (subtab: SubTab) {
     if (subtab.action) {
        actionHandlers[subtab.action]?.();
    }
}

function handlePaneClick(pos: XYPosition) {
    console.log("PANEL CLIDK", pos);
}

async function onFileSelected(event: Event) {
  loading.value = true
  const file = (event.target as HTMLInputElement).files?.[0]
  if (!file) return

  const reader = new FileReader()
  reader.onload = async (e) => {
    const base64 = e.target?.result as string
    backgroundImage.value = base64
    await SaveMapBackground(base64, file.name)

    const img = new Image()
    img.onload = () => {
      nodes.value = nodes.value.filter(n => n.type !== 'background')
      nodes.value.unshift({
        id: 'background',
        type: 'background',
        position: { x: 0, y: 0 },
        data: { src: base64, width: img.naturalWidth, height: img.naturalHeight },
        draggable: false,
        selectable: false,
        connectable: false,
      })
      loading.value = false
    }
    img.src = base64
  }
  reader.readAsDataURL(file)
  ;(event.target as HTMLInputElement).value = ''
}

async function createNode() {
    try {
    const x = 10
    const y = 10
    const label = 'New Node'
    const nodeId = await SaveNode(label, x, y)
    
    nodes.value.push({
      id: String(nodeId),
      type: 'room',
      position: { x, y },
      data: { label, type: 'room' },
    } satisfies FlowNode)
  } catch (e) {
    logger.error('failed to create node', { error: String(e) })
  }
}

async function onNodeDragStop(node: VueFlowNode) {
  await UpdateNode(Number(node.id), node.data.label, node.position.x, node.position.y)
}

async function FillNodes() {
  const tmpNodes = await GetAllNodes()
  nodes.value = tmpNodes.map(n => ({
    id: String(n.Id),
    type: n.Type,
    position: { x: n.Position.X, y: n.Position.Y },
    data: { label: n.Data.Label, type: n.Data.Type as RoomData['type'] },
    draggable: n.Draggable ?? undefined,
    selectable: n.Selectable ?? undefined,
    connectable: n.Connectable ?? undefined,
  } satisfies FlowNode))
}

onMounted(async () => {
  try {
    await FillNodes()
    const saved = await GetMapBackground()
    if (saved) {
      backgroundImage.value = saved
      const img = new Image()
      img.onload = () => {
        nodes.value.unshift({
          id: 'background',
          type: 'background',
          position: { x: 0, y: 0 },
          data: { src: saved, width: img.naturalWidth, height: img.naturalHeight },
          draggable: false,
          selectable: false,
          connectable: false,
        })
      }
      img.src = saved
    }
  } catch (e) {
    logger.error('failed to load map background', { error: String(e) })
  }
})

provide('updateNodeLabel', async (id: string, label: string, x: number, y: number) => {
  try {
    await UpdateNode(Number(id), label, x, y)
  } catch (e) {
    logger.error('failed to update node', { error: String(e) })
  }
})

</script>
<template>
    <div class="map-editor">
        <h1>Map Editor</h1>
         <Tabs :tabs="tabs" @sub-tab-change="subTabChangeHandler"/>
         <NodeDeletePanel
          v-if="showDeletePanel"
          v-model:selected-node-id="selectedNodeId"
          :nodes="deletableNodes"
          @delete="deleteSelectedNode"
    />
          <div class="canvas-container">
            <FlowCanvas 
            :background-image="backgroundImage"
            :nodes="nodes"
            :edges="edges"
             @on-node-drag-stop="onNodeDragStop"
             />
         </div>
         <input 
            ref="fileInput" 
            type="file"
            accept="image/png, image/jpeg, image/webp"
            style="display:none"
            @change="onFileSelected"
            />
        <div v-if="loading" class="spinner" />
    </div>
</template>
<style>
.map-editor {
    display: flex;
    flex-direction: column;
    flex: 1;
    min-height: 0;
}

.canvas-container {
  padding: 10px;
  flex: 1;
  position: relative;
  min-height: 0;
  padding: 15px;
}

.spinner {
  width: 32px; height: 32px;
  border: 3px solid #ddd;
  border-top-color: #6366f1;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}
@keyframes spin { to { transform: rotate(360deg); } }
</style>