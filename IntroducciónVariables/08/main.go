package main

import "fmt"

func main() {
	var a complex128 = 3 + 7.6i
	b := real(a)
	c := imag(a)

	fmt.Printf("La variable compleja es de tipo %T.\n", a)
	fmt.Printf("La parte real es de tipo %T.\n", b)
	fmt.Printf("La parte imaginaria es de tipo %T.\n", c)
}
