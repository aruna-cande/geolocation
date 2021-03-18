FROM golang:1.16

RUN echo "deb http://packages.cloud.google.com/apt gcsfuse-jessie main" | tee /etc/apt/sources.list.d/gcsfuse.list \
    && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - \
    && apt-get update \
    && apt-get -y install gcsfuse

RUN mkdir importer
COPY . /importer
WORKDIR /importer/cmd/importer
RUN go build -o ./app/importer .
#ENTRYPOINT ["./app/importer"]
#CMD ["/tmp/data/data_dump.csv"]