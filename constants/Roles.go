package main
import "fmt"

const (
	isAdmin= 1 << iota
	isManager
	isMember
)

func main(){

	var roles byte = isAdmin | isManager | isMember
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is admin? %v\n", isAdmin & roles == isAdmin)
}