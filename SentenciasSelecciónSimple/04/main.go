package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Dime un número entero: ")
	fmt.Scan(&numero)

	if (numero % 2) == 0 {
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}

}
