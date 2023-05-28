package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		inicial float64
		redito  float64
		tiempo  float64
		final   float64
	)

	fmt.Print("Dime la cantidad inicial: ")
	fmt.Scan(&inicial)
	fmt.Print("Dime el rédito en tantos por ciento: ")
	fmt.Scan(&redito)
	fmt.Print("Dime el tiempo: ")
	fmt.Scan(&tiempo)

	final = inicial * math.Pow(1+redito/100, tiempo)
	fmt.Println("La candidad final es:", final)

}
