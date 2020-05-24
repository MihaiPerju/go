package main
import "fmt"

func main(){
	 
	a := []int{}
	
	// Variatic function. From 2nd value all of the values are getting appended
	a = append(a, 2,3,4,5,6)
	
	fmt.Printf("Size: %v\n",len(a))
	fmt.Printf("Capacity: %v\n",cap(a))
}