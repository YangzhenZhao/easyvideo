import { useStore } from 'vuex'

export const bytesFormat = (bytesSize: number) => {
  const units = ['B', 'KB', 'MB', 'GB', 'TB']
  for (let i = 0; i < 4; i++) {
    if (bytesSize < 1024) {
      return `${bytesSize.toFixed(2)} ${units[i]}`
    }
    bytesSize /= 1024
  }
}

export const serverAddress = () => {
  const store = useStore()
  if (store.state.serverAddress === '') {
    const currentLocation = window.location
    const address = `${currentLocation.protocol}//${currentLocation.hostname}:8000`
    store.commit('setServerAddress', address)
  }
  return store.state.serverAddress
}
