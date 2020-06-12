package main

import "fmt"

func main() {
	res, err := divide(4, 0)
	if err {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("Cannot divide by 0")
	}
	return a / b, nil
}
