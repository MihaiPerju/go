package main

import "fmt"

func main() {
	var s *int = sum(1, 2)
	fmt.Println(*s)
}

func sum(a, b int) *int {
	var s int = a + b
	return &s
}
