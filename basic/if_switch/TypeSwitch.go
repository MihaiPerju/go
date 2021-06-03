package main

import "fmt"

func main() {

	var i interface{} = 4

	switch i.(type) {
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("Other")
	}
}
