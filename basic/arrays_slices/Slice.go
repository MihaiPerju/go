package main
import "fmt"

func main(){
	a:=[]int{1,2,3,4,5}

	// Referencing works like a pointer
	b:=a
	b[0]=99

	slice := a[1:2]

	fmt.Println(a)
	fmt.Println(slice)
}