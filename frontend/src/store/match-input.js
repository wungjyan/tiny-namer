import { defineStore } from 'pinia'

export const useMatchInputStore = defineStore('matchInput', () => {
  const pattern = ref('')
  const replacement = ref('')

  function setPattern(newValue) {
    pattern.value = newValue
  }

  function setReplacement(newValue) {
    replacement.value = newValue
  }

  return {
    pattern,
    setPattern,
    replacement,
    setReplacement,
  }
})
