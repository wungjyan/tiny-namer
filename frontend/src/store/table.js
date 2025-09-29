import { defineStore } from 'pinia'
import { computed } from 'vue'

export const useTableStore = defineStore('table', () => {
  const data = ref([])

  const hasUpdate = computed(() => {
    return data.value.some(item => item.success)
  })

  function setData(newData) {
    data.value = newData
  }

  function removeTableRow(index) {
    data.value.splice(index, 1)
  }

  return { data, hasUpdate, setData, removeTableRow }
})
