package main

import (
	"fmt"
)

var a int
var b string
var c bool
var d float64

func main() {
	a = 10
	b = "Batman"
	c = true
	d = 20.99

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

	// fmt package session

	// Interpreted strings
	x := "Oi tudo bem?\ncomo vai?\tespero que bem"

	fmt.Println(x)

	// Raw strings
	y := `"Oi tudo bem?\ncomo vai?\tespero que bem"`

	fmt.Println(y)
}
