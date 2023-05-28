package main

import "fmt"

func main() {
	var (
		a float64
		b float64
	)

	fmt.Print("Dime el primer número: ")
	fmt.Scan(&a)
	fmt.Print("Dime el segundo número: ")
	fmt.Scan(&b)

	fmt.Println("La suma es:", a+b)
	fmt.Println("La resta es:", a-b)
	fmt.Println("El producto es:", a*b)
	fmt.Println("La división es:", a/b)
}
