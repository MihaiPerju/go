package main

import "fmt"

func main() {

	sum(1, 2, 3, 4)
}

func sum(values ...int) {

	var sum int = 0
	for _, v := range values {
		sum += v
	}

	fmt.Println(sum)

}
