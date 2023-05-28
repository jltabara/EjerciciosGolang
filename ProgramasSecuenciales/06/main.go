package main

import "fmt"

func main() {
	var (
		base      float64
		altura    float64
		area      float64
		perimetro float64
	)

	fmt.Print("Dime la base: ")
	fmt.Scan(&base)
	fmt.Print("Dime la altura: ")
	fmt.Scan(&altura)
	area = base * altura
	perimetro = 2*base + 2*altura
	fmt.Println("El área es:", area)
	fmt.Println("El perímetro es:", perimetro)
}
