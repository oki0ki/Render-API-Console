FROM node:20-alpine AS svelte-build
WORKDIR /app
COPY console/package*.json ./
RUN npm ci
COPY console/ ./
RUN npm run build

FROM golang:1.22-alpine AS go-build
WORKDIR /app
COPY *.go go.mod ./
RUN go build -o gateway .

FROM nginx:alpine
RUN apk add --no-cache bash
COPY --from=svelte-build /app/build /usr/share/nginx/html
COPY --from=go-build /app/gateway /app/gateway
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY start.sh /start.sh
RUN chmod +x /start.sh
EXPOSE 7860
CMD ["/start.sh"]
