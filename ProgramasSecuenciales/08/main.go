package main

import "fmt"

func main() {
	const c = 299792
	var distancia float64

	distancia = c * 60 * 60 * 24 * 365 * 1000

	fmt.Printf("La distancia de la Tierra al Sol son %f metros", distancia)
}
