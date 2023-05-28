package main

import "fmt"

func main() {
	var (
		precio   float64
		cantidad int
	)
	fmt.Print("Dime el precio del artículo: ")
	fmt.Scan(&precio)
	fmt.Print("Dime la cantidad de artículos: ")
	fmt.Scan(&cantidad)
	if cantidad > 3 {
		fmt.Println("El precio total es", float64(cantidad)*precio*0.8)
	} else {
		fmt.Println("El precio total es", float64(cantidad)*precio)
	}

}
