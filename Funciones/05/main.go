package main

import "fmt"

func main() {
	a, b := nombreApellido()
	fmt.Println("Mi nombre es", a, "y mi apellido es", b)
}

func nombreApellido() (string, string) {
	return "José Luis", "Tábara"
}
