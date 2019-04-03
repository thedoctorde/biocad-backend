FROM golang:1.11

WORKDIR $GOPATH/src/biocad

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["biocad"]