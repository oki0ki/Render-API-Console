#!/bin/sh
export PORT=${PORT:-10000}

PORT=3000 node /app/svelte/build &
SVELTE_PID=$!

( wait $SVELTE_PID; echo "SvelteKit exited, shutting down container"; kill 1 ) &

exec /app/gateway
