package main

import "fmt"

func main() {
	var a = [7]int64{1, 1, 2, 3, 5, 8, 13}
	fmt.Println("El array es:", a)
	var b = []int64{1, 1, 2, 3, 5, 8, 13}
	fmt.Println("El slice es:", b)
}
