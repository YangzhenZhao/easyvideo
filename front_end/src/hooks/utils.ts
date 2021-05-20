export const bytesFormat = (bytesSize: number) => {
  const units = ['B', 'KB', 'MB', 'GB', 'TB']
  for (let i = 0; i < 4; i++) {
    if (bytesSize < 1024) {
      return `${bytesSize.toFixed(2)} ${units[i]}`
    }
    bytesSize /= 1024
  }
}
