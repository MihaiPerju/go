package main
import "fmt"

const(
	_ = iota // ignore the first zero value
	KB  = 1 << (10*iota)
	MB
	GB
	TB
)
func main(){
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
}