FROM golang:alpine AS builder
ARG APP_NAME=echoserver
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ${APP_NAME}
FROM alpine:latest
COPY --from=builder /app/${APP_NAME} /app/
WORKDIR /app

CMD ["./echoserver"]