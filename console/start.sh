#!/bin/sh
export PORT=${PORT:-10000}

PORT=3000 node /app/svelte/build &

exec /app/gateway
