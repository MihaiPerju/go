package main

import "fmt"

type person struct {
	age int
}

func main() {

	var michael *person

	michael = new(person)
	// one way of dereferencing pointer structs
	(*michael).age = 22

	// Another way of dereferencing
	fmt.Println(michael.age)

}
