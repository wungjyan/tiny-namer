import { defineStore } from 'pinia'

export const useMatchSettingsStore = defineStore('matchSettings', () => {
  // 匹配设置
  const ignoreExt = ref(true)
  const ignoreCase = ref(true)
  const global = ref(true)

  function setIgnoreExt(newValue) {
    ignoreExt.value = newValue
  }

  function setIgnoreCase(newValue) {
    ignoreCase.value = newValue
  }

  function setGlobal(newValue) {
    global.value = newValue
  }

  function setAllSettings(arr) {
    ignoreExt.value = arr.includes('ignoreExt')
    ignoreCase.value = arr.includes('ignoreCase')
    global.value = arr.includes('global')
  }

  return {
    ignoreExt,
    ignoreCase,
    global,
    setIgnoreExt,
    setIgnoreCase,
    setGlobal,
    setAllSettings,
  }
})
