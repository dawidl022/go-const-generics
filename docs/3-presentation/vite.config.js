import { defineConfig } from 'vite'
import hmr from './hmr'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    hmr()
  ]
})
