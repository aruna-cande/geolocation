FROM golang:1.16

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY pkg .
WORKDIR /app
RUN ls -a
RUN go  build -o app ./cmd/geolocation-api
CMD ["/app"]