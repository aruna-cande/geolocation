FROM golang:1.16

RUN mkdir importer
COPY . /importer
WORKDIR /importer/cmd/importer
RUN go build -o ./app/importer .
ENTRYPOINT ["./app/importer"]
CMD ["data_dump.csv"]