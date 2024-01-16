package main

import (
	"fmt"
)

var foraDoCodeBlock = "Olá bom dia dentro do code block"

func main() {
	dentroDoCodeBlock := "Olá bom dia dentro do code block"

	fmt.Println(dentroDoCodeBlock)
	fmt.Println(foraDoCodeBlock)

	x := 12
	y := "String"
	z := true

	fmt.Println(x, y, z)

	x = 20
	y = "Hello"
	z = false

	fmt.Println(x, y, z)

	_, numero2 := fmt.Println(200, 100)
	fmt.Println(numero2)
}
