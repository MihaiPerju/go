FROM golang

WORKDIR /go/src/go-fiber-crm-basic

COPY database /go/src/go-fiber-crm-basic/database
COPY lead /go/src/go-fiber-crm-basic/lead
COPY main.go .

RUN ls -l

RUN go mod init github.com/MihaiPerju/go-fiber-crm-basic

RUN go get "github.com/gofiber/fiber"
RUN go get "github.com/jinzhu/gorm"
RUN go get "github.com/jinzhu/gorm/dialects/sqlite"

# RUN go build
