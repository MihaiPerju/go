package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName:  "Anderson"}
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact.email = "admin@app.com"
	alex.contact.zipCode = 45544

	(&alex).updateName("Alecs")
	alex.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p.firstName)
}
