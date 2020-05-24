package main
import "fmt"

func main(){
	
	studentGrades := map[string]int{
		"Michael": 9,
		"Dan": 7,
		"Eve": 10,
	}

	statePopulations := make(map[string]int)
	statePopulations["Moldova"] = 4000000

	delete(studentGrades,"Dan")

	// ok is false if the value is not present in the struct
	grade,ok := studentGrades["Dan"]
	fmt.Println(grade,ok)

	fmt.Println(studentGrades["Michael"])
	fmt.Println(statePopulations)
}