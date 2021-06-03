package main

import "fmt"

func main() {
	n := 4
	switch i := n + 2; i {
	case 1, 5, 10:
		fmt.Println(1, 5, 10)
	case 2, 4, 6:
		fmt.Println(2, 4, 6)
	default:
		fmt.Println("Other")
	}
}
