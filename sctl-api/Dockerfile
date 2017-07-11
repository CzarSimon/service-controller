FROM golang:1.8-alpine
RUN apk update && apk upgrade && apk add --no-cache git gcc musl-dev

RUN mkdir -p /usr/src/sctl-api
COPY . /usr/src/sctl-api

WORKDIR /usr/src/sctl-api

RUN go get github.com/mattn/go-sqlite3
RUN go get github.com/CzarSimon/sctl-common
RUN go get github.com/CzarSimon/util
RUN go install github.com/mattn/go-sqlite3
RUN export GOPATH=$GOPATH:$PWD
RUN go build

EXPOSE 9104
CMD ["./sctl-api"]
