# syntax=docker/dockerfile:1
FROM golang:1.21-alpine AS build

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build the binary
COPY . .
RUN go build -o oneart-identity-service ./cmd/main.go

# Minimal image for production
FROM alpine:3.17
WORKDIR /app

COPY --from=build /app/oneart-identity-service /app/oneart-identity-service

EXPOSE 8080

CMD ["/app/oneart-identity-service"]