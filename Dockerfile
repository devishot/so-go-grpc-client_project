FROM golang:1.11-alpine as builder

RUN apk --update add git openssh && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

WORKDIR /go/src/github.com/devishot/so-go-grpc-client_project

COPY . .

RUN go get ./

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o app" -command="./app"

EXPOSE 9001