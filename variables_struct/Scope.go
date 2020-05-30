package main
import "fmt"

// Uppercase variable name makes a variable being a global one
var A int = 0

// Lowercase variables are scoped to the package
var b int = 1

func main(){
	// Block scope - only available inside the main function
	var c int = 2

	fmt.Printf("A is global variable with value %v\n", A)
	fmt.Printf("b is package variable with value %v\n", b)
	fmt.Printf("c is block variable with value %v\n", c)

}