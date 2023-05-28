package main

import "fmt"

func main() {
	x, y := parametroTipo(4, 7, 8, 9)
	fmt.Println("El valor es:", x, "y el tipo es:", y)
}

func parametroTipo(n ...int) ([]int, string) {
	return n, fmt.Sprintf("%T", n)
}
