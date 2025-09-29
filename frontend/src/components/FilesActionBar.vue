<template>
  <n-space justify="start" align="center">
    <n-button @click="chooseFiles">
      <template #icon>
        <n-icon>
          <FileTraySharp />
        </n-icon>
      </template>
      选择文件
    </n-button>
    <n-button v-if="tableData.length" @click="clearFiles">
      <template #icon>
        <n-icon>
          <CloseCircleSharp />
        </n-icon>
      </template>
      清空选择
    </n-button>
  </n-space>
</template>

<script setup>
  import { FileTraySharp, CloseCircleSharp } from '@vicons/ionicons5'
  import { SelectFiles } from '@go'
  import { useTableStore } from '@/store'

  const tableStore = useTableStore()
  const tableData = computed(() => tableStore.data)
  async function chooseFiles() {
    const res = await SelectFiles()
    console.log('文件选择结果==', res)
    if (res.success) {
      tableStore.setData(res.files)
    }
  }

  function clearFiles() {
    tableStore.setData([])
  }
</script>
