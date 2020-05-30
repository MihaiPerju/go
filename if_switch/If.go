package main

import "fmt"

func main() {
	n := 20
	if n == 30 {
		fmt.Println("OK")
	} else if n < 30 {
		fmt.Println("Less than 30")
	} else {
		fmt.Println("Greater than 30")
	}
}
