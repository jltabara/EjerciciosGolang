package main

import "fmt"

func main() {
	var a [7]int64

	for i := 0; i < len(a); i++ {
		fmt.Print("Dime el valor ", i, ": ")
		fmt.Scanln(&a[i])
	}
	fmt.Println(a)
}
