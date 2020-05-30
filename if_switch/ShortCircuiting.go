package main

import "fmt"

func main() {
	if 20 < 3 && isTrue() {
		fmt.Println("There's a problem here")
	} else {
		fmt.Println("Done.")
	}
}

func isTrue() bool {
	fmt.Println("Returning true")
	return true
}
