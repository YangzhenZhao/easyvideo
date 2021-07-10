import { createStore } from 'vuex'

let serverAddress = ''
if (localStorage.serverAddress) {
  serverAddress = localStorage.serverAddress
}
const store = createStore({
  state: {
    serverAddress: serverAddress
  },
  mutations: {
    setServerAddress (state, serverAddress) {
      localStorage.serverAddress = serverAddress
      state.serverAddress = serverAddress
    }
  },
  actions: {
  }
})

export default store
