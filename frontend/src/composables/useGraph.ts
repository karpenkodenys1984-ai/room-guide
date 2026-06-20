import { ref } from 'vue'
import { useVueFlow, type Connection } from '@vue-flow/core'
import type { FlowNode, FlowEdge } from '../types/graph'

export function useGraph() {
 const nodes = ref<FlowNode[]>([])
 
  const edges = ref<FlowEdge[]>([])

  function addNode(x: number, y: number, label = 'Новый узел') {
    const id = crypto.randomUUID()
    nodes.value.push({
      id,
      type: 'room',
      position: { x, y },
      data: { label, type: 'room' },
    })
  }

  function onConnect(connection: Connection) {
    edges.value.push({
      id: `${connection.source}-${connection.target}`,
      source: connection.source!,
      target: connection.target!,
    })
  }

  function serialize() {
    return JSON.stringify({ nodes: nodes.value, edges: edges.value })
  }

  function load(json: string) {
    const data = JSON.parse(json)
    nodes.value = data.nodes
    edges.value = data.edges
  }

  return { nodes, edges, addNode, onConnect, serialize, load }
}