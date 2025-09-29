<template>
  <n-layout has-sider class="h-[calc(100vh-40px)]">
    <n-layout-sider bordered show-trigger collapse-mode="width" :collapsed-width="64" :width="180">
      <n-menu
        v-model:value="activeMenuKey"
        :collapsed-width="64"
        :collapsed-icon-size="22"
        :options="menuOptions"
      />
    </n-layout-sider>
    <n-layout-content>
      <n-split
        direction="horizontal"
        style="height: 100%"
        :max="0.45"
        :min="0.25"
        :default-size="0.35"
      >
        <template #1>
          <template v-if="currentMode">
            <n-card title="命名设置" embedded :bordered="false" size="small">
              <component :is="currentMode" />
            </n-card>
            <n-card embedded :bordered="false" size="small" class="mt-10">
              <n-collapse>
                <n-collapse-item title="查看帮助" name="1">
                  <p>这里是帮助说明</p>
                </n-collapse-item>
              </n-collapse>
            </n-card>
          </template>
        </template>
        <template #resize-trigger>
          <n-el style="width: 0; height: 100%; border-right: 1px solid var(--divider-color)"></n-el>
        </template>
        <template #2>
          <n-card title="" :bordered="false" size="small">
            <FilesActionBar class="mb-4" />
            <FilesTable />
          </n-card>
        </template>
      </n-split>
    </n-layout-content>
  </n-layout>
</template>

<script setup>
  import {
    CodeSlashOutline,
    TextOutline,
    PencilOutline,
    SwapVerticalOutline,
    CheckmarkDoneOutline,
    ExtensionPuzzleOutline,
    TrashOutline,
  } from '@vicons/ionicons5'
  import { NIcon } from 'naive-ui'
  import { Regex, TextReplace } from './components/mode/index'
  import FilesTable from './components/FilesTable.vue'
  import FilesActionBar from './components/FilesActionBar.vue'
  import { shallowRef } from 'vue'

  function renderIcon(icon) {
    return () => h(NIcon, null, { default: () => h(icon) })
  }
  const menuOptions = [
    { label: '正则表达式', key: 'regex', icon: renderIcon(CodeSlashOutline) },
    { label: '文本替换', key: 'textReplace', icon: renderIcon(TextOutline) },
    { label: '插入文本', key: 'insertText', icon: renderIcon(PencilOutline) },
    { label: '大小写转换', key: 'caseConversion', icon: renderIcon(SwapVerticalOutline) },
    { label: '添加序号', key: 'addSerial', icon: renderIcon(CheckmarkDoneOutline) },
    { label: '扩展名修改', key: 'extEdit', icon: renderIcon(ExtensionPuzzleOutline) },
    { label: '移除字符', key: 'removeChar', icon: renderIcon(TrashOutline) },
  ]

  const modeMap = {
    regex: Regex,
    textReplace: TextReplace,
  }
  const currentMode = shallowRef(Regex)

  const activeMenuKey = ref(menuOptions[0].key)

  watch(
    () => activeMenuKey.value,
    val => {
      currentMode.value = modeMap[val]
    }
  )
</script>
