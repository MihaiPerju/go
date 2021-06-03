package main
import "fmt"

type Doctor struct{
	name string
	id int
	patients [3]string
}

func main(){
	michael := Doctor{
		name: "Michael",
		id:4,
		patients: [3]string{"Daniel","John","Boris"},
	}
	daniel := michael
	daniel.name = "Daniel"
	
	fmt.Println(michael)
}