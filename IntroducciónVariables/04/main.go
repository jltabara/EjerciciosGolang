package main

import "fmt"

func main() {
	var (
		a8  int8  = 13
		a16 int16 = 23
		a32 int32 = 42
		a64 int64 = 234
	)

	fmt.Println("El contenido de las variables es:",
		a8, a16, a32, a64)
}
