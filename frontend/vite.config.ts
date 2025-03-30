import { defineConfig } from 'vite'
import { resolve } from 'path'
import react from '@vitejs/plugin-react-swc'

const root = resolve(__dirname, 'src')

export default defineConfig({
  plugins: [react()],
  root,
  publicDir: resolve(__dirname, 'public'),
  build: {
    outDir: resolve(__dirname, 'dist'),
    rollupOptions: {
      input: {
        top: resolve(root, 'index.html'),
        planningpoker: resolve(root, 'planning-poker', 'index.html'),
      }
    }
  }
})
