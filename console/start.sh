#!/bin/sh
API_PORT=3001 /app/gateway &
nginx -g 'daemon off;'
