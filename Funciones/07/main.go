package main

import "fmt"

func main() {
	a := 89
	b := esPar(a)
	fmt.Println("El número", a, "es par:", b)
}

func esPar(numero int) bool {
	return (numero % 2) == 0
}
