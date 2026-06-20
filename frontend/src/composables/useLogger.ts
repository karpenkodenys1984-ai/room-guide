import { getCurrentInstance } from 'vue'
import { createLogger } from '../lib/logger'

export function useLogger() {
  const instance = getCurrentInstance()
  const name = instance?.type.__name ?? instance?.type.name ?? 'Unknown'
  return createLogger(name)
}