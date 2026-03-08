# syntax=docker/dockerfile:1

# -- Build stage --
FROM golang:1.23-alpine AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app ./cmd/importer

# -- gcsfuse install stage --
FROM golang:1.23-alpine AS gcsfuse-builder

RUN apk add --no-cache git fuse
RUN go install github.com/googlecloudplatform/gcsfuse@latest

# -- Runtime stage --
FROM alpine:3.20

RUN apk add --no-cache ca-certificates fuse \
    && addgroup -S app && adduser -S app -G app

COPY --from=gcsfuse-builder /go/bin/gcsfuse /usr/local/bin/gcsfuse

RUN mkdir -p /tmp/data && chown app:app /tmp/data

COPY --from=builder /app /app

USER app
ENTRYPOINT ["/app"]