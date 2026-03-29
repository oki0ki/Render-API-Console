import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
        plugins: [sveltekit()],
        server: {
                host: '0.0.0.0',
                port: parseInt(process.env.PORT || '5173'),
                allowedHosts: true,
                hmr: {
                        protocol: 'wss',
                        clientPort: 443,
                        host: process.env.REPLIT_DEV_DOMAIN || 'localhost',
                },
                proxy: {
                        '/v1': {
                                target: 'http://localhost:3001',
                                changeOrigin: true,
                        }
                }
        },
        build: {
                target: 'esnext',
                sourcemap: false,
                minify: 'esbuild',
                rollupOptions: {
                        output: {
                                manualChunks: undefined
                        }
                }
        }
});
