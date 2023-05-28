package main

import "fmt"

func main() {
	var (
		a float64 = 9.56
		b float64 = 8.234
	)

	fmt.Println("La suma es:", a+b)
	fmt.Println("La resta es:", a-b)
	fmt.Println("El producto es:", a*b)
	fmt.Println("La división es:", a/b)
}
