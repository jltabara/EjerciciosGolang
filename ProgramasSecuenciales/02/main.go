package main

import "fmt"

func main() {
	var (
		a int64
		b int64
	)

	fmt.Print("Dime el primer número: ")
	fmt.Scan(&a)
	fmt.Print("Dime el segundo número: ")
	fmt.Scan(&b)

	fmt.Println("La suma es:", a+b)
	fmt.Println("La resta es:", a-b)
	fmt.Println("El producto es:", a*b)
	fmt.Println("El cociente es:", a/b)
	fmt.Println("El resto es:", a%b)
}
