package main

import "fmt"

type person struct {
	age  int
	name string
}

// Operating on a COPY - pure function. The function can also use a pointer and mutate
func (p person) identify() {
	fmt.Println("Name", p.name)
	fmt.Println("Age", p.age)
}

func main() {
	michael := new(person)
	michael.age = 22
	michael.name = "perju"

	michael.identify()
}
