import { createStore } from 'vuex'

let serverAddress = 'http://127.0.0.1:5000'
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
