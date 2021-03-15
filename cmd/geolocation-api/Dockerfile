FROM golang:1.16
EXPOSE 8080

RUN mkdir geolocation
COPY . /geolocation
WORKDIR /geolocation/cmd/geolocation-api
RUN go build -o /app .
ENTRYPOINT ["/app"]