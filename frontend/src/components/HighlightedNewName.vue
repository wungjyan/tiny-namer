<template>
  <n-highlight
    :text="newBaseName || baseName"
    :patterns="[replacement]"
    :highlight-style="`background-color: ${themeVars.errorColor};color: white;`"
  />
  <span v-if="ignoreExt">{{ ext }}</span>
  <n-highlight
    v-else
    :text="newExt || ext"
    :patterns="[replacement]"
    :highlight-style="`background-color: ${themeVars.errorColor};color: white;`"
  />
</template>

<script setup>
  import { useMatchInputStore, useMatchSettingsStore } from '@/store'
  import { useThemeVars } from 'naive-ui'
  const themeVars = useThemeVars()

  const matchInputStore = useMatchInputStore()
  const matchSettingsStore = useMatchSettingsStore()
  const ignoreExt = computed(() => matchSettingsStore.ignoreExt)
  const replacement = computed(() => matchInputStore.replacement)

  const props = defineProps({
    baseName: {
      type: String,
      default: '',
    },
    newBaseName: {
      type: String,
      required: true,
    },
    newExt: {
      type: String,
      default: '',
    },
    ext: {
      type: String,
      default: '',
    },
  })
</script>
