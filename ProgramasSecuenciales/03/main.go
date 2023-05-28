package main

import "fmt"

func main() {
	var (
		precio    float64
		precioIVA float64
	)

	fmt.Print("Dime el precio del producto sin IVA: ")
	fmt.Scan(&precio)
	precioIVA = precio * 1.21
	fmt.Println("El precio con IVA es:", precioIVA)
}
