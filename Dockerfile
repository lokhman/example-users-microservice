FROM golang:alpine

COPY . /go/src/github.com/lokhman/example-users-microservice
WORKDIR /go/src/github.com/lokhman/example-users-microservice

RUN apk update && \
    apk add --no-cache git gcc musl-dev && \
    go get -u github.com/codegangsta/gin && \
    go get -u github.com/golang/dep/cmd/dep && \
    go get -u github.com/swaggo/swag/cmd/swag && \
    dep ensure && \
    swag init

CMD gin -i run
