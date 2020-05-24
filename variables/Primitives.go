package main
import "fmt"

func main(){

	// Boolean
	var n bool = false
	fmt.Printf("%v %T\n", n, n)


	// Numeric values
	var a int8
	var b int16
	var c int32
	var d int64
	// Anything unitiliased in Go is 0/false

	fmt.Println(a,b,c,d)


	// Binary operators
	a1:=10
	b1:=3

	fmt.Println(a1&b1)
	fmt.Println(a1|b1)
	fmt.Println(a1^b1)
	fmt.Println(a1&^b1)

	// Bit shifting
	fmt.Println(a1<<3)
	fmt.Println(a1>>3)

	// Floating points
	var o float32 = 3.14
	var m float64 = 4.18

	fmt.Println(o,m)

	// Complex types
	var c1 complex64 = 1+2i
	var c2 complex128 = 3+4i

	fmt.Println(complex128(c1)+c2)
	// Real and imaginary part
	fmt.Println(real(c1),imag(c2))


	// Text data types
	var s string = "ok"
	fmt.Printf("%v %T\n",s[0],s[0])
	// Concatenation
	fmt.Println(s+s)

	// Converting a string to bytes
	bb := []byte(s)
	fmt.Println(bb)

	// rune
	r:='a'
	var rr rune = 'b'
	fmt.Printf("%v %T\n",r,r)
	fmt.Printf("%v %T\n",rr,rr)

}