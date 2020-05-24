package main
import "fmt"

type Person struct{
	age int
	name string
}

type Student struct{
	Person
	studentId int
}

func main(){
	michael := Student{
		Person : Person{
			name : "Michael",
			age : 21,
		},
		studentId : 1009,
	}

	michael.age = 22

	fmt.Println(michael)
}