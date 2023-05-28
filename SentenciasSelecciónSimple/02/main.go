package main

import "fmt"

func main() {
	var edad int
	fmt.Print("Dime tu edad: ")
	fmt.Scan(&edad)
	if edad >= 18 {
		fmt.Println("Eres mayor de edad.")
	} else {
		fmt.Println("Eres menor de edad.")
	}

}
