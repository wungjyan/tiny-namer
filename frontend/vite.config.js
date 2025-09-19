import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import AutoImport from "unplugin-auto-import/vite";
import { NaiveUiResolver } from "unplugin-vue-components/resolvers";
import Components from "unplugin-vue-components/vite";
import tailwindcss from "@tailwindcss/vite";

const rootPath = new URL(".", import.meta.url).pathname;

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    tailwindcss(),
    AutoImport({
      imports: [
        "vue",
        {
          "naive-ui": ["useThemeVars"],
        },
      ],
    }),
    Components({
      resolvers: [NaiveUiResolver()],
    }),
  ],
  resolve: {
    alias: {
      "@": rootPath + "src",
      wailsjs: rootPath + "wailsjs",
    },
  },
});
