import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import svgr from 'vite-plugin-svgr';
import { env } from 'process';
import { openSync } from 'fs';
import { resolve } from 'path';
import path from 'path';

const commitHash = env.GIT_COMMIT || 'local';
const version = env.VERSION || 'local';

// This plugin creates a keep file to include the
// dist directory to the version control but exclude the content.
const keep = {
  name: 'Create static keep file for git',
  closeBundle() {
    openSync(resolve(__dirname, 'dist/keep'), 'w');
  },
};

// https://vitejs.dev/config/
export default defineConfig({
  define: {
    __COMMIT_HASH__: JSON.stringify(commitHash),
    __VERSION__: JSON.stringify(version),
  },
  server: {
    proxy: {
      '/api': 'http://localhost:8080',
    },
  },
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
  plugins: [react(), svgr(), keep],
});
