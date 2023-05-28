package main

import "fmt"

func main() {
	a := suma(4, 7, 8, 9)
	fmt.Println("La suma es", a)
}

func suma(ns ...int) int {
	s := 0
	for _, v := range ns {
		s += v
	}
	return s
}
