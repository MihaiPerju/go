package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error", err)
		}
	}()
	panic("Panic happening")
}
