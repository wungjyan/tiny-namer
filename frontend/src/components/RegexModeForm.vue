<template>
  <n-form ref="formRef" :model="model" :rules="rules" label-placement="left">
    <n-form-item label="匹配规则：" path="input1">
      <n-input v-model:value="model.input1" placeholder="请输入正则表达式" />
    </n-form-item>
    <n-form-item label="替换内容：" path="input2">
      <n-input v-model:value="model.input2" placeholder="请输入替换内容" />
    </n-form-item>
    <n-form-item label="匹配设置：">
      <n-checkbox-group v-model:value="model.checkbox">
        <n-checkbox value="ignoreExt">忽略扩展名</n-checkbox>
        <n-checkbox value="ignoreCase">忽略大小写</n-checkbox>
        <n-checkbox value="global">全局匹配</n-checkbox>
      </n-checkbox-group>
    </n-form-item>
    <n-form-item>
      <n-button type="primary" @click="handlePreview">命名预览</n-button>
    </n-form-item>
  </n-form>
</template>

<script setup>
import { ref } from "vue";
import { useTemplateRef } from "vue";
import { ValidateRegex } from "wailsjs/go/main/App";

const model = ref({
  input1: "",
  input2: "",
  checkbox: [],
});

const rules = {
  input1: [
    {
      required: true,
      validator: async (rule, value) => {
        if (value.trim() === "") {
          return Promise.reject(new Error("请输入正则表达式"));
        }

        try {
          const result = await ValidateRegex(value);
          if (!result.valid) {
            return Promise.reject(new Error(result.message));
          }
          return Promise.resolve();
        } catch (error) {
          return Promise.reject(new Error("正则表达式校验失败"));
        }
      },
      trigger: ["blur"],
    },
  ],
};

const formRef = useTemplateRef("formRef");

function handlePreview(e) {
  e.preventDefault();

  console.log("开始表单校验，当前值:", model.value);

  formRef.value?.validate((errors) => {
    console.log("校验结果 - errors:", errors);
    if (!errors) {
      console.log("表单校验通过", model.value);
      // 这里可以调用预览功能
    } else {
      console.log("表单校验失败", errors);
    }
  });
}
</script>
