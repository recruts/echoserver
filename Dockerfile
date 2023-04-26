FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o echoserver
FROM alpine:latest
COPY --from=builder /app/echoserver /app/
WORKDIR /app

CMD ["./echoserver"]