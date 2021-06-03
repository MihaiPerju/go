package main

import "fmt"

func main() {

	// Function signature
	var divide func(float64, float64) (float64, error)

	// Function assignment
	divide = func(a, b float64) (float64, error) {
		if b == 0 {
			return 0.0, fmt.Errorf("Can't divide by 0")
		}
		return a / b, nil

	}

	res, err := divide(3, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}
