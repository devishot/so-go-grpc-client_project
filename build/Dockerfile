# First stage
FROM golang:1.11-alpine as builder

WORKDIR /go/src/github.com/devishot/so-go-grpc-client_project
COPY . .

RUN apk --update add git openssh && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

RUN go get ./
RUN go build -o app

# Second stage
FROM alpine:latest

WORKDIR /www-data

COPY --from=builder /go/src/github.com/devishot/so-go-grpc-client_project/app .

CMD ["./app"]

EXPOSE 9001