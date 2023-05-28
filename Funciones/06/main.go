package main

import "fmt"

func main() {
	a := 90
	b := 12
	c, d := cocienteResto(a, b)
	fmt.Printf("El cociente de %v y %v es %v y el resto es %v.\n", a, b, c, d)
}

func cocienteResto(a int, b int) (int, int) {
	return a / b, a % b
}
