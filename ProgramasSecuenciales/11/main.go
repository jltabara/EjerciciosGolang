package main

import "fmt"

func main() {
	var (
		practica      float64
		teoria        float64
		participacion float64
		notaFinal     float64
	)

	fmt.Print("Dime la nota de práctica: ")
	fmt.Scan(&practica)
	fmt.Print("Dime la nota de teoría: ")
	fmt.Scan(&teoria)
	fmt.Print("Dime la nota de participación: ")
	fmt.Scan(&participacion)

	notaFinal = 0.3*practica + 0.6*teoria + 0.1*participacion

	fmt.Println("La nota final es:", notaFinal)
}
