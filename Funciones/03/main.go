package main

import "fmt"

func main() {
	a := 42
	b := doble(a)
	fmt.Println("El doble de", a, "es", b)
}

func doble(n int) int {
	return 2 * n
}
