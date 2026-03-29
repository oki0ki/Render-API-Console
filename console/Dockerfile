FROM node:20-alpine AS svelte-build
WORKDIR /app
COPY console/console/package*.json ./
RUN npm install
COPY console/console/ ./
RUN npm run build

FROM golang:1.22-alpine AS go-build
WORKDIR /app
COPY console/*.go console/go.mod ./
RUN go build -o gateway .

FROM node:20-alpine
RUN apk add --no-cache nginx gettext bash
WORKDIR /app
COPY --from=svelte-build /app/build ./svelte/build
COPY --from=svelte-build /app/node_modules ./svelte/node_modules
COPY --from=svelte-build /app/package.json ./svelte/package.json
COPY --from=go-build /app/gateway ./gateway
COPY console/nginx.conf /etc/nginx/nginx-template.conf
COPY console/start.sh /start.sh
RUN chmod +x /start.sh
EXPOSE 10000
CMD ["/start.sh"]
