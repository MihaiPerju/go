package main

import "fmt"

func main() {

	i := 10
	switch {
	case i <= 10:
		fmt.Println(1)
		fallthrough
	case i <= 20:
		fmt.Println(2)
	default:
		fmt.Println("Other")
	}
}
