<template>
  <n-config-provider :theme="theme">
    <n-split
      direction="horizontal"
      :default-size="0.3"
      :min="0.3"
      :max="0.75"
      style="height: 100vh"
    >
      <!-- 左侧操作区 -->
      <template #1>
        <div class="h-full pt-10">
          <n-scrollbar class="max-h-full px-4">
            <div class="flex flex-col items-center gap-5">
              <!-- 文件导入 -->
              <n-button type="primary" size="large" @click="chooseFiles"
                ><template #icon>
                  <NIcon>
                    <FileTrayIcon />
                  </NIcon> </template
                >选择文件</n-button
              >

              <!-- 重命名方式 -->
              <n-card title="重命名方式" size="small" embedded>
                <n-select
                  v-model:value="renameMode"
                  :options="modeOptions"
                  placeholder="选择重命名方式"
                />
              </n-card>

              <!-- 动态配置表单（示例） -->
              <n-card v-if="modeComponent" embedded>
                <component :is="modeComponent" />
              </n-card>
            </div>
          </n-scrollbar>
        </div>
      </template>

      <!-- 右侧文件列表 -->
      <template #2>
        <div class="h-full p-4">
          <n-data-table
            :columns="columns"
            :data="fileList"
            :bordered="false"
            size="small"
          />
        </div>
      </template>
    </n-split>
    <n-global-style />
  </n-config-provider>
</template>

<script setup>
import { FileTray as FileTrayIcon } from "@vicons/ionicons5";
import { computed, ref } from "vue";
import RegexModeForm from "./components/RegexModeForm.vue";
import TextModeForm from "./components/TextModeForm.vue";

const renameMode = ref(null);
const modeComponent = computed(() => {
  switch (renameMode.value) {
    case "regex":
      return RegexModeForm;
    case "text":
      return TextModeForm;
    default:
      return null;
  }
});

const modeOptions = [
  { label: "正则表达式", value: "regex" },
  { label: "文本替换", value: "text" },
  { label: "插入文本", value: "insert" },
  { label: "大小写转换", value: "case" },
  { label: "添加序号", value: "serial" },
  { label: "扩展名修改", value: "ext" },
  { label: "移除字符", value: "remove" },
];

// 预览表格数据
const fileList = ref([
  {
    name: "example1.txt",
    newName: "example1.txt",
    path: "/folder",
    ext: "txt",
  },
  {
    name: "example2.jpg",
    newName: "example2.jpg",
    path: "/folder",
    ext: "jpg",
  },
]);

const columns = [
  { title: "原文件名", key: "name" },
  { title: "预览新文件名", key: "newName" },
  { title: "所在路径", key: "path" },
  { title: "扩展名", key: "ext" },
];

const chooseFiles = () => {
  // TODO: 调用 wails runtime.OpenMultipleFilesDialog
};

const chooseFolder = () => {
  // TODO: 调用 wails runtime.OpenDirectoryDialog
};

const preview = () => {
  // TODO: 生成预览结果
};

const apply = () => {
  // TODO: 执行重命名
};
</script>

<style scoped>
/* 简单布局 */
</style>
