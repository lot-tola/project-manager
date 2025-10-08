import { defineConfig } from "vite";
import path from "node:path";
import vue from "@vitejs/plugin-vue";
import tailwind from "@tailwindcss/vite";

export default defineConfig({
  plugins: [vue(), tailwind()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
  server: {
    port: 5173,
  },
});
