package main

import "fmt"

// Global vairables have to have the type declared
var a int = 42;

func main(){
	// Declaring the variable
	var i int = 4
	fmt.Println(i)

	// Declaring a variable, but let go automatically decide its type
	auto := 4
	fmt.Printf("%v %T\n", auto, auto)


	// Printing the variable and its type
	j := 2.
	fmt.Printf("%v %T\n",j,j)

	// Printing global variables
	fmt.Println(a)

}
