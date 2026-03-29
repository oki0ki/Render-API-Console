#!/bin/sh
export PORT=${PORT:-10000}

envsubst '${PORT}' < /etc/nginx/nginx-template.conf > /etc/nginx/nginx.conf

API_PORT=3001 /app/gateway &
PORT=3000 node /app/svelte/build &

nginx -g 'daemon off;'
