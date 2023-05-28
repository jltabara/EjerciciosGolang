package main

import "fmt"

func main() {
	var (
		a int64 = 995
		b int64 = 8
	)

	fmt.Println("La suma es:", a+b)
	fmt.Println("La resta es:", a-b)
	fmt.Println("El producto es:", a*b)
	fmt.Println("El cociente es:", a/b)
	fmt.Println("El resto es:", a%b)
}
