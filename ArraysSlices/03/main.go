package main

import "fmt"

func main() {
	var a [5]int64 = [5]int64{4, 6, 7, 8, 9}
	var b []int64 = []int64{4, 7, 8, 9, 4}
	fmt.Printf("El tipo del array es %T\n", a)
	fmt.Printf("El tipo del slice es %T\n", b)
}
