package main

import "fmt"

func main() {
	s := []int{5, 6, 7}

	for k, v := range s {
		fmt.Println(k, v)
	}

	v := map[string]int{
		"michael": 10,
		"daniel":  13,
		"john":    19,
	}

	for k, v := range v {
		fmt.Println(k, v)
	}

	str := "hello world"
	for k, v := range str {
		fmt.Println(k, v)
	}
}
