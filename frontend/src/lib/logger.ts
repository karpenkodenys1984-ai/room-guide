import { LogFromFrontend } from '../../wailsjs/go/logger/Bridge'
import { logger } from '../../wailsjs/go/models'

export function createLogger(component: string) {
  function send(level: string, message: string, extra?: Record<string, unknown>) {
    const fn = level === 'error' ? console.error
             : level === 'warn'  ? console.warn
             : level === 'debug' ? console.debug
             : console.log
    fn(`[${component}]`, message, extra ?? '')

    if (level === 'error' || level === 'warn') {
      const entry = logger.FrontendLogEntry.createFrom({
        level,
        message,
        component,
        extra: extra ?? {},
        timestamp: new Date(),
      })

      LogFromFrontend(entry)
        .catch((err) => console.error('log bridge failed', err))
    }
  }

  return {
    debug: (msg: string, extra?: Record<string, unknown>) => send('debug', msg, extra),
    info:  (msg: string, extra?: Record<string, unknown>) => send('info',  msg, extra),
    warn:  (msg: string, extra?: Record<string, unknown>) => send('warn',  msg, extra),
    error: (msg: string, extra?: Record<string, unknown>) => send('error', msg, extra),
  }
}