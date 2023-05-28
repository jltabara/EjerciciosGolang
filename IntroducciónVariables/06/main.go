package main

import "fmt"

func main() {
	var (
		a64  complex64  = 4 + 9.6i
		a128 complex128 = complex(2.5, 7.34)
	)

	fmt.Println("El contenido de las variables es:",
		a64, a128)
}
