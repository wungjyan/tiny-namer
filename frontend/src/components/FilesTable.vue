<template>
  <n-data-table :columns="columns" :data="tableData" :bordered="false" size="medium">
    <template #empty>
      <n-empty description="未选择文件"></n-empty>
    </template>
  </n-data-table>
  <div v-if="tableData.length" class="mt-5">
    <n-button type="primary" :disabled="!tableStore.hasUpdate">应用新名称</n-button>
  </div>
</template>

<script setup>
  import HighlightedOriginalName from './HighlightedOriginalName.vue'
  import { NTag, NButton } from 'naive-ui'
  import { useTableStore } from '@/store'
  const tableStore = useTableStore()
  const tableData = computed(() => tableStore.data)

  const columns = [
    {
      title: '原文件名',
      key: 'fullName',
      render(row) {
        return h(HighlightedOriginalName, {
          fileName: row.fullName,
          baseName: row.baseName,
          ext: row.ext,
        })
      },
    },
    {
      title: '预览新文件名',
      render(row) {
        return h('span', {}, row.changed ? row.newFullName : '-')
      },
    },
    {
      title: '状态',
      render(row) {
        return row.changed
          ? h(
              NTag,
              {
                type: 'info',
                bordered: false,
              },
              { default: () => '已改变' }
            )
          : h('span', {}, '-')
      },
    },
    {
      title: '操作',
      align: 'center',
      render(_, index) {
        return h(
          NButton,
          {
            quaternary: true,
            size: 'tiny',
            onClick: () => tableStore.removeTableRow(index),
          },
          {
            default: () => '移除',
          }
        )
      },
    },
  ]
</script>
