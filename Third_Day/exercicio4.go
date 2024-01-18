package main

import (
	"fmt"
)

type hamburguer int

var c int

var a hamburguer

func exercicio4() {
	// Exercicio 4 parte 1
	fmt.Printf("%v\t", a)
	fmt.Printf("%T\t", a)

	a = 42

	fmt.Printf("%v\t", a)

	// Exercicio 4 parte 2 ou exercicio 5
	c = int(a)
	fmt.Println(c)
	fmt.Printf("%T\t", c)
}
