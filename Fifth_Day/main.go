package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := "é"

	a := "e"

	c := "#99"

	fmt.Printf("%v, %v, %v\n", a ,x ,c)

	f := []byte(x)
	d := []byte(a)
	e := []byte(c)

	fmt.Printf("%v, %v, %v\n", f ,d ,e)

	g := 10
	y := 10.0

	//Valores defaults, quando vc não diz ao compilador o que você quer
	// Output: 10, float64 e 10, int
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", g, g)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	var i uint16 = 65535
	fmt.Println(i)
	i++
	fmt.Println(i)
}