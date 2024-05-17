package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	user := pessoa{"Vinicius", "Pimentel", 12, 21}
	fmt.Println(user)

	studant := estudante{user, "Ciencia da computacao", "UAM"}
	fmt.Println(studant.nome)
}
