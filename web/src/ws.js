import RobustWebSocket from 'robust-websocket'
import config from './config'
import vm from './main'

class WS {
  constructor () {
    this.socket = null
    this.handlers = {}
  }

  close = () => {
    this.socket.close()
  }

  init = () => {
    if (this.socket) {
      return this
    }

    this.socket = new RobustWebSocket(`${config.wsURL}`)

    this.socket.onopen = () => {
      vm.$store.commit('SetWSStatus', { status: 1 })
      this.socket.send(JSON.stringify({
        command: 'init'
      }))
    }

    this.socket.onclose = () => {
      vm.$store.commit('SetWSStatus', { status: 3 })
    }

    this.socket.onmessage = event => {
      const payload = JSON.parse(event.data) || {}
      const command = payload.command
      if (typeof this.handlers[command] === 'function') {
        this.handlers[command]({
          payload
        })
      }
    }

    this.send = this.socket.send

    return this
  }

  on = (command, handler) => {
    this.handlers[command] = handler
  }
}

export default new WS()
