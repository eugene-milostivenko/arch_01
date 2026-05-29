# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Зависимостей нет, но копируем go.mod для корректной сборки модуля
COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o temperature-api

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/temperature-api .

# Порт по умолчанию из задания
EXPOSE 8081

CMD ["./temperature-api"]
