package main

import (
	"fmt"
)

type hotdog int

var b hotdog = 10

func main() {
	x := 10
	fmt.Printf("%T\n", x)
	fmt.Printf("%T", b)

	x = int(b)
	fmt.Printf("%v\n", x)

	// Isso não funciona!!!
	// b = x cannot use x (variable of type int) as hotdog Psalue in assignment

	exercicio1()

	exercicio2()

	exercicio4()
}
