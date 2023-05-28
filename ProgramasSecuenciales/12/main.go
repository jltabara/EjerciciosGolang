package main

import "fmt"

func main() {
	var (
		a string = "Hola"
		b string = "Mundo"
		c string = a + b
	)
	fmt.Println("El contenido de la variable es:", c)
}
