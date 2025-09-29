<template>
  <n-config-provider :theme="theme">
    <n-message-provider>
      <div class="h-100vh">
        <n-el
          tag="div"
          class="h-[40px] box-border flex justify-end items-center pr-5"
          style="border-bottom: 1px solid var(--border-color); --wails-draggable: drag"
        ></n-el>
        <AppContent />
      </div>
    </n-message-provider>
    <n-global-style />
  </n-config-provider>
</template>

<script setup>
  import AppContent from './AppContent.vue'
  import { darkTheme, useOsTheme } from 'naive-ui'
  const osTheme = useOsTheme()
  const theme = computed(() => (osTheme.value === 'dark' ? darkTheme : null))
</script>
