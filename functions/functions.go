package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	result := somar(9, 20)
	fmt.Println(result)

	var f = func(txt string) string {
		fmt.Println("Funcao f")
		return txt
	}

	fmt.Println(f("testando"))
	resultadoSoma, resultadoSubtracao := calculos(10, 5)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}
