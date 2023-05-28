package main

import "fmt"

func main() {
	var (
		x int64
		y int64
	)
	fmt.Print("Dime el primer número entero: ")
	fmt.Scan(&x)
	fmt.Print("Dime el segundo número entero: ")
	fmt.Scan(&y)
	fmt.Println("Antes del cambio los valores son:", x, y)
	x, y = y, x
	fmt.Println("Después del cambio los valores son:", x, y)
}
