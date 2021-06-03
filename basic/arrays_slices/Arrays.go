package main
import "fmt"

func main(){
	// ... means create an array that will pick up the right size
	grades := [...]int{7,8,9}
	var students [3]string

	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of students: %v\n", len(students))
}