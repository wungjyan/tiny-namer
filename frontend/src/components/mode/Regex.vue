<template>
  <n-form ref="formRef" :model="model" :rules="rules" label-placement="left" class="w-full">
    <n-form-item label="匹配规则：" path="pattern">
      <n-input
        v-model:value="model.pattern"
        placeholder="输入正则表达式"
        @update:value="handleUpdatePattern"
        clearable
      />
    </n-form-item>

    <n-form-item label="替换内容：" path="replacement">
      <n-input
        v-model:value="model.replacement"
        placeholder="输入替换内容，不填则替换为空"
        @update:value="handleUpdateReplacement"
        clearable
      />
    </n-form-item>
    <n-form-item label="匹配设置：">
      <n-checkbox-group v-model:value="model.settings" @update:value="handleUpdateChecked">
        <n-checkbox value="ignoreExt">忽略扩展名</n-checkbox>
        <n-checkbox value="ignoreCase">忽略大小写</n-checkbox>
        <n-checkbox value="global">全局匹配</n-checkbox>
      </n-checkbox-group>
    </n-form-item>
  </n-form>
  <n-button type="primary" @click="handlePreview">
    {{ tableStore.hasUpdate ? '重新生成预览' : '生成预览' }}
  </n-button>
</template>

<script setup>
  defineOptions({
    name: 'Regex',
  })
  import { useMessage } from 'naive-ui'
  import { useTableStore, useMatchSettingsStore, useMatchInputStore } from '@/store'
  import { ValidateRegex, PreviewRegexRename } from '@go'

  const message = useMessage()

  const tableStore = useTableStore()
  const matchSettingsStore = useMatchSettingsStore()
  const matchInputStore = useMatchInputStore()

  const tableData = computed(() => tableStore.data)

  const model = ref({
    pattern: '',
    replacement: '',
    settings: ['ignoreExt', 'ignoreCase', 'global'],
  })

  function handleUpdatePattern() {
    matchInputStore.setPattern(model.value.pattern)
    if (tableStore.hasUpdate) {
      previewRename()
    }
  }
  function handleUpdateReplacement() {
    matchInputStore.setReplacement(model.value.replacement)
    if (tableStore.hasUpdate) {
      previewRename()
    }
  }

  function handleUpdateChecked() {
    matchSettingsStore.setAllSettings(model.value.settings)
    if (tableStore.hasUpdate) {
      previewRename()
    }
  }

  const rules = {
    pattern: [
      {
        validator: async (rule, value) => {
          try {
            const result = await ValidateRegex(value)
            if (!result) {
              return Promise.reject(new Error('正则表达式无效'))
            }
            return Promise.resolve()
          } catch (error) {
            return Promise.reject(new Error('正则表达式校验失败'))
          }
        },
        trigger: ['blur'],
      },
    ],
  }

  const formRef = useTemplateRef('formRef')

  function handlePreview(e) {
    if (!tableData.value.length) {
      message.warning('需要先选择文件才可以预览')
      return
    }
    if (!model.value.pattern) {
      message.warning('没有输入匹配规则，不会产生预览')
      if (tableStore.hasUpdate) {
        previewRename()
      }
      return
    }
    e.preventDefault()
    formRef.value?.validate(errors => {
      if (!errors) {
        // 这里可以调用预览功能
        previewRename()
      } else {
        console.log('表单校验失败', errors)
      }
    })
  }

  async function previewRename() {
    const req = {
      files: tableData.value,
      pattern: model.value.pattern,
      replacement: model.value.replacement,
      ignoreExt: matchSettingsStore.ignoreExt,
      ignoreCase: matchSettingsStore.ignoreCase,
      global: matchSettingsStore.global,
    }
    const res = await PreviewRegexRename(req)
    console.log('预览结果===', res)
    if (res.success) {
      tableStore.setData(res.items)
    }
  }
</script>
