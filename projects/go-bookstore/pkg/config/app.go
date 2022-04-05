package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)


var (
	db *gorm.DB
)

func Connect(){
	DBMS := "mysql"

	USER := "michael"
	PASS := "thisIsThePassword"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "simplerest"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true"

    d, err := gorm.Open(DBMS, CONNECT)

	if err!=nil{
		panic(err)	
	}

	fmt.Println("Seems like a connection has been established")

	db = d
}

func GetDB() *gorm.DB{
	return db
}