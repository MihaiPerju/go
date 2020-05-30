package main
import (
	"fmt"
	"strconv"
)

func main(){

	i:=4
	var j float32 = float32(i)
	// converting an integer to a string
	var char string = strconv.Itoa(i)

	fmt.Println(j)
	fmt.Printf("%v %T\n", char, char)
}