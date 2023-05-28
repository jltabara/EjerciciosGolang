package main

import "fmt"

func main() {
	var (
		peso      float64
		pesoMarte float64
	)
	fmt.Print("Dime tu peso en Tierra: ")
	fmt.Scan(&peso)
	pesoMarte = peso * 0.38
	fmt.Println("Tu peso en Marte es:", pesoMarte)
}
