package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func exercicio2() {
	// Exercicio 2 parte 1
	fmt.Printf("%v\n%v\n%v\n", x, y, z)

	fmt.Println("O nome do valor atribuído as variáveis se chama valor zero")

	// Exercicio 2 parte 2 ou exercicio 3
	s := fmt.Sprintf("\nEu tenho %d e me chamo %s, isso é verdade? %v", x, y, z)

	fmt.Println(s)
}
