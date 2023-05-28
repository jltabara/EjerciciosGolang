package main

import "fmt"

func main() {
	var numero int64
	var unidades int64

	fmt.Print("Dime un número entero positivo: ")
	fmt.Scan(&numero)

	unidades = numero % 10

	fmt.Println("La cifra de las unidades es:", unidades)
}
