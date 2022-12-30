import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import svgr from "vite-plugin-svgr";
import { env } from "process";

const commitHash = env.GIT_COMMIT || "local";
const version = env.VERSION || "local";

// https://vitejs.dev/config/
export default defineConfig({
  define: {
    __COMMIT_HASH__: JSON.stringify(commitHash),
    __VERSION__: JSON.stringify(version),
  },
  server: {
    proxy: {
      "/api": "http://localhost:8080",
    },
  },
  plugins: [
    react(),
    svgr(),
  ],
});
