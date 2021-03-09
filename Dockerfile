FROM golang:1.16

RUN mkdir /internal
COPY /internal/app/geolocation-api /internal
WORKDIR /internal/geolocation-api

RUN go build .