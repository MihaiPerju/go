FROM golang

WORKDIR /go/src/go-fiber-mongo-hrms

COPY main.go .
COPY go.mod .
COPY go.sum .

# RUN go env -w GO111MODULE=off
RUN go get "github.com/gofiber/fiber/v2"
RUN go get "go.mongodb.org/mongo-driver/mongo/options"
RUN go get "go.mongodb.org/mongo-driver/mongo"
RUN go get "go.mongodb.org/mongo-driver/bson"
RUN go get "go.mongodb.org/mongo-driver/bson/primitive"

RUN go build