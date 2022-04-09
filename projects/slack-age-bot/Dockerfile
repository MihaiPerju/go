FROM golang

WORKDIR /go/src/slack-age-bot

RUN go env -w GO111MODULE=off
RUN go get "github.com/gorilla/mux"
RUN go get "github.com/shomali11/slacker"