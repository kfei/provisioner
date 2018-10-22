import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    available: 0,
    wsStatus: 0
  },

  mutations: {
    SetAvailable (state, { count = 0 }) {
      state.available = count
    },

    SetWSStatus (state, { status = 0 }) {
      state.wsStatus = status
    }
  }
})
