package main

import "fmt"

func main() {
	var (
		pizzas           int64
		personas         int64
		porcionesPersona int64
		resto            int64
	)

	fmt.Print("Dime el número de pizzas: ")
	fmt.Scan(&pizzas)
	fmt.Print("Dime el número de personas: ")
	fmt.Scan(&personas)

	porcionesPersona = (pizzas * 8) / personas
	resto = (pizzas * 8) % personas
	fmt.Println("Cada persona come", porcionesPersona, "porciones.")
	fmt.Println("Además han sobrado", resto, "porciones.")
}
