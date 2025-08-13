# Сборка
FROM golang:1.23.8 AS builder
WORKDIR /app
COPY ./go.mod ./
RUN go mod download
COPY ./ /app
RUN GO111MODULE=auto CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -o todotask cmd/main.go

# Запуск
FROM alpine:3.20.3
RUN apk add tzdata
WORKDIR /app
COPY --from=builder /app/todotask .
ENTRYPOINT [ "./todotask" ]