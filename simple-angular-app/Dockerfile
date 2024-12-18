FROM node:23-alpine3.19 AS builder

WORKDIR /app
COPY . .

RUN npm install
RUN npm run build

FROM nginx:1.27.3-alpine

RUN rm -rf /usr/share/nginx/html/*
COPY --from=builder /app/dist/simple-angular-app/browser /usr/share/nginx/html
CMD \
  mainFiles=$(ls /usr/share/nginx/html/main*.js) \
  && \
  for f in ${mainFiles}; do \
    envsubst '$API_HOST,$API_PORT' < "$f" > "${f}.tmp" && mv "${f}.tmp" "$f"; \
  done \
  && \
  nginx -g 'daemon off;'
EXPOSE 80
