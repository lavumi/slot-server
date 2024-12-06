# Build stage
FROM golang:1.19-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build both servers
RUN CGO_ENABLED=0 GOOS=linux go build -o build/web-server cmd/web.go
RUN CGO_ENABLED=0 GOOS=linux go build -o build/slot-server cmd/slot.go

# Web server stage
FROM alpine:latest AS web-server
WORKDIR /
COPY --from=builder /app/web ./web
COPY --from=builder /app/build/web-server ./
RUN apk add --no-cache bash
EXPOSE 8081
CMD ["/web-server"]

# Slot server stage
FROM alpine:latest AS slot-server
WORKDIR /
COPY --from=builder /app/parSheet ./parSheet
COPY --from=builder /app/build/slot-server ./
RUN apk add --no-cache bash
EXPOSE 8088
CMD ["/slot-server"]