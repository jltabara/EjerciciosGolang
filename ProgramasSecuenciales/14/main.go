package main

import "fmt"

func main() {
	var (
		euros   float64
		pesetas float64
	)
	const cambio = 166.386

	fmt.Print("Dime una cantidad de euros: ")
	fmt.Scan(&euros)

	pesetas = euros * cambio

	fmt.Println(euros, "euros son ", pesetas, "pesetas")
}
