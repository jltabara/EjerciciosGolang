package main

import "fmt"

func main() {
	var (
		lado      float64
		area      float64
		perimetro float64
	)
	fmt.Print("Dime el lado del cuadrado: ")
	fmt.Scan(&lado)
	area = lado * lado
	perimetro = 4 * lado
	fmt.Println("El área es:", area)
	fmt.Println("El perímetro es:", perimetro)

}
