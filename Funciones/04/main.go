package main

import "fmt"

func main() {
	a := 4
	b := 8
	c := sumar(a, b)
	fmt.Println("La suma de", a, "y", b, "es", c)
}

func sumar(a, b int) int {
	return a + b
}
