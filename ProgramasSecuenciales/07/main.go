package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		radio    float64
		longitud float64
		area     float64
	)

	fmt.Print("Dime el radio: ")
	fmt.Scan(&radio)
	longitud = 2 * math.Pi * radio
	area = math.Pi * radio * radio
	fmt.Println("La longitud es:", longitud)
	fmt.Println("El área es:", area)
}
