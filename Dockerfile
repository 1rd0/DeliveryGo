# Используем официальный образ Golang
FROM golang:1.23 AS builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . .
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/main.go

# Используем минимальный образ для продакшена (убираем лишние зависимости)
FROM alpine:latest
WORKDIR /
COPY --from=builder /server /server
COPY --from=builder /app/config /config
RUN chmod +x /server
EXPOSE 50051
CMD ["/server"]