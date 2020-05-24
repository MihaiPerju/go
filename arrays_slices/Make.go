package main
import "fmt"

func main(){
	a := make([]int, 3, 100) //type, size, capacity
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("Size: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

}