package main

import "fmt"

func main() {
	var (
		num1 int
		num2 int
	)
	fmt.Print("Dime el primer número: ")
	fmt.Scan(&num1)
	fmt.Print("Dime el segundo número: ")
	fmt.Scan(&num2)
	if (num2 % num1) == 0 {
		fmt.Println(num1, "es divisor de", num2)
	} else {
		fmt.Println(num1, "no es divisor de", num2)
	}

}
