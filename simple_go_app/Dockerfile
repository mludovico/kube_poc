FROM golang:1.22.1-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o todo_api

FROM alpine:3.20.3

COPY --from=builder /app/todo_api /todo_api
COPY ./data /data
EXPOSE $API_PORT
ENTRYPOINT [ "/todo_api" ]