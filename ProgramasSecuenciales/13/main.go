package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		c1 float64
		c2 float64
		h  float64
	)

	fmt.Print("Dime el primer cateto: ")
	fmt.Scan(&c1)
	fmt.Print("Dime el segundo cateto: ")
	fmt.Scan(&c2)

	h = math.Sqrt(c1*c1 + c2*c2)

	fmt.Println("La hipotenusa mide:", h)
}
