export interface RoomData {
  label: string
  type: 'room' | 'entrance' | 'exit'
}

export interface BackgroundData {
  src: string
  width: number
  height: number
}

export type FlowNode = {
  id: string
  type: string
  position: { x: number; y: number }
  data: RoomData | BackgroundData
  draggable?: boolean
  selectable?: boolean
  connectable?: boolean
}

export type FlowEdge = {
  id: string
  source: string
  target: string
}