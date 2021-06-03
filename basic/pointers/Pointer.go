package main

import "fmt"

func main() {
	var a int = 4
	var b *int = &a
	*b = 8

	fmt.Println(a)
}
