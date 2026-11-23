import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// The Go backend has no CORS middleware, so in dev we proxy /api -> :8080
// and strip the prefix. That way the browser only ever talks to one origin.
export default defineConfig({
  plugins: [react()],
  server: {
    port: 5173,
    proxy: {
      '/api': {
        target: process.env.BACKEND_URL ?? 'http://localhost:8080',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
});
