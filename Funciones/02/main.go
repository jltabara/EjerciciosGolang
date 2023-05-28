package main

import "fmt"

func main() {
	saludar("Pepe")
}

func saludar(nombre string) {
	fmt.Println("Hola", nombre, "desde una función.")
}
