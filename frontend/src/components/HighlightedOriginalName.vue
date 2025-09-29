<template>
  <n-highlight :text="baseName" :patterns="highlightedMatches" />
  <span v-if="ignoreExt">{{ ext }}</span>
  <n-highlight v-else :text="ext" :patterns="highlightedMatches" />
</template>

<script setup>
  import { GetHighlightParts } from '@go'
  import { useMatchInputStore, useMatchSettingsStore } from '@/store'

  const props = defineProps({
    fileName: {
      type: String,
      required: true,
    },
    baseName: {
      type: String,
      default: '',
    },
    ext: {
      type: String,
      default: '',
    },
  })

  const matchInputStore = useMatchInputStore()
  const matchSettingsStore = useMatchSettingsStore()

  const ignoreExt = computed(() => {
    return matchSettingsStore.ignoreExt
  })

  // 存储匹配到的文本片段
  const highlightedMatches = ref([])

  // 监听属性变化，调用Go后端方法获取高亮数据
  const updateHighlights = async () => {
    if (!matchInputStore.pattern) {
      highlightedMatches.value = []
      return
    }

    try {
      const matches = await GetHighlightParts({
        fileName: props.fileName,
        baseName: props.baseName,
        pattern: matchInputStore.pattern,
        ignoreExt: matchSettingsStore.ignoreExt,
        ignoreCase: matchSettingsStore.ignoreCase,
        global: matchSettingsStore.global,
      })
      console.log('高亮词汇===', matches)

      highlightedMatches.value = matches || []
    } catch (error) {
      // 如果调用失败，清空匹配结果
      highlightedMatches.value = []
    }
  }

  // 监听所有相关属性的变化
  watch(
    [
      () => props.fileName,
      () => props.baseName,
      () => matchInputStore.pattern,
      () => matchSettingsStore.ignoreExt,
      () => matchSettingsStore.ignoreCase,
      () => matchSettingsStore.global,
    ],
    updateHighlights,
    { immediate: true }
  )
</script>
