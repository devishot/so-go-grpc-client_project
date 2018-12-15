FROM golang:1.11-alpine

COPY . /go/src/github.com/devishot/so-go-grpc-client_project/
WORKDIR /go/src/github.com/devishot/so-go-grpc-client_project

RUN go get ./
RUN go build -o app

CMD ["./app"]

EXPOSE 8080