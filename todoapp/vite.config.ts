import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import path from 'path'

export default defineConfig({
  server: {    // <-- this object is added
    port: 5656
  },
  plugins: [svelte()],
  resolve: {
    alias: {
      $root: path.resolve('./src'),
    },
  },
})
