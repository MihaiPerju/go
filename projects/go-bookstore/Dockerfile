FROM golang

WORKDIR /go/src/go-bookstore


COPY cmd /go/src/go-bookstore/cmd
COPY pkg /go/src/go-bookstore/pkg

RUN go mod init github.com/MihaiPerju/go-bookstore

RUN go get "github.com/jinzhu/gorm"
RUN go get "github.com/jinzhu/gorm/dialects/mysql"
RUN go get "github.com/gorilla/mux"


WORKDIR /go/src/go-bookstore/cmd
RUN go build
