package main

import "fmt"

func main() {
	a := "Pepe"
	b := saludar(a)
	fmt.Printf("La función ha escrito \"%v\".\n", b)
}

func saludar(nombre string) string {
	return fmt.Sprintf("Hola %v", nombre)
}
