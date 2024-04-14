import path, { resolve } from 'path';

import { defineConfig } from 'vite';
import { env } from 'process';
import { openSync } from 'fs';
/// <reference types="vitest" />
import react from '@vitejs/plugin-react';
import svgr from 'vite-plugin-svgr';

const commitHash = env.GIT_COMMIT || 'local';
const version = env.VERSION || 'local';

// This plugin creates a keep file to include the
// dist directory to the version control but exclude the content.
const keep = {
  closeBundle() {
    openSync(resolve(__dirname, 'dist/keep'), 'w');
  },
  name: 'Create static keep file for git',
};

// https://vitejs.dev/config/
export default defineConfig({
  define: {
    __COMMIT_HASH__: JSON.stringify(commitHash),
    __VERSION__: JSON.stringify(version),
  },
  plugins: [react(), svgr(), keep],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
    },
  },
  server: {
    proxy: {
      '/api': 'http://localhost:8080',
    },
  },
  test: {
    coverage: {
      provider: 'istanbul',
    },
    environment: 'jsdom',
    setupFiles: './src/setupTests.ts',
  },
});
