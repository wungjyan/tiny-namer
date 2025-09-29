<template>
  <n-select
    v-model:value="renameMode"
    :options="modeOptions"
    placeholder="选择重命名方式"
    @update:value="handleUpdateValue"
  />
</template>

<script setup>
  const renameMode = ref(null)

  const modeOptions = [
    { label: '正则表达式', value: 'regex' },
    { label: '文本替换', value: 'text' },
    { label: '插入文本', value: 'insert' },
    { label: '大小写转换', value: 'case' },
    { label: '添加序号', value: 'serial' },
    { label: '扩展名修改', value: 'ext' },
    { label: '移除字符', value: 'remove' },
  ]

  const emit = defineEmits(['change'])

  function handleUpdateValue(val) {
    emit('change', val)
  }
</script>
