package main

import (
	"fmt"
)

const oi string = "Oi, tudo bem?"
const x int64 = 32

const (
	a = 10
	b = 20
	d = 30
	z = iota
	// Underscore is a blank identifier, "throw away" value
	_ = iota
)

var c int64

func main() {
	// Print the decimal, binary, and hex values of 42
	fmt.Printf("%d\t%b\t%#x\n", 42, 42, 42)

	fmt.Println(oi)

	a := 10.0
	b := 10

	c = x

	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	fmt.Println(a, b, d)

	fmt.Println(z)

	// Bit shifting
	x := 2

	y := x << 1

	fmt.Printf("%b\n", y)

	// Exercises
	exercicio1()

	exercicio2()

	exercicio3()

	exercicio4()

}
