# syntax=docker/dockerfile:1

# -- Build stage --
FROM golang:1.23-alpine AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app ./cmd/geolocation-api

# -- Runtime stage --
FROM alpine:3.20

RUN apk add --no-cache ca-certificates \
    && addgroup -S app && adduser -S app -G app

COPY --from=builder /app /app

USER app
EXPOSE 8080
ENTRYPOINT ["/app"]