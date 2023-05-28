package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Dime un número entero: ")
	fmt.Scan(&numero)
	if numero > 0 {
		fmt.Println("El número es positivo.")
	} else {
		fmt.Println("El número no es positivo.")
	}

}
