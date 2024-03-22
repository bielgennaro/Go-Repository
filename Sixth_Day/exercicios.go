package main

import (
	"fmt"
)

const tipada string = "Constante tipada"
const naoTipada = "Constante não tipada"

func exercicio1() {
	fmt.Printf("%d\t%b\t%#x\n", 42, 42, 42)
}

func exercicio2() {
	x := 42 == 43
	y := 42 <= 43
	z := 42 > 43
	b := 42 >= 43
	c := 42 < 43

	fmt.Println(x, y, z, b, c)

}

func exercicio3() {
	fmt.Println(tipada)
	fmt.Println(naoTipada)
}

func exercicio4() {
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	b := a << 1

	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
}
