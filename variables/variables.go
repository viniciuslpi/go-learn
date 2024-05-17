package main

import (
	"errors"
	"fmt"
)

func main() {

	var (
		variable1 = "teste1"
		variable2 = "teste2"
	)

	variable3 := "inferencia de tipo"
	variable4, variable5 := "variable4", "variable5"

	fmt.Println(variable1, variable2, variable3)

	const constante2 string = "constante"

	fmt.Println(constante2)

	variable5, variable4 = variable4, variable5

	fmt.Println(variable4, variable5)

	var erro error = errors.New("Erro interno")

	fmt.Println(erro)

}
