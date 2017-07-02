FROM golang:1.8-alpine
RUN apk update && apk upgrade && apk add --no-cache git nodejs

RUN mkdir -p /usr/src/sctl-api
COPY . /usr/src/sctl-api

WORKDIR /usr/src/sctl-api

RUN go install github.com/mattn/go-sqlite3
RUN go get github.com/CzarSimon/util
RUN export GOPATH=$GOPATH:$PWD
RUN go build

EXPOSE 1321
CMD ["./sctl-api"]
