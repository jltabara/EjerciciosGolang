package main

import "fmt"

func main() {
	const c = 299792
	var distancia float64
	var tiempo float64

	fmt.Print("Dime la distancia en km: ")
	fmt.Scan(&distancia)
	tiempo = distancia / c
	fmt.Println("La luz tarda", tiempo, "segundos")
}
