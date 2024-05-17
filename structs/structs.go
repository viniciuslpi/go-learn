package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var user usuario
	user.nome = "Vinicius"
	user.idade = 23
	enderecoTeste := endereco{"Rua marmelada", 32}
	user.endereco = enderecoTeste

	user2 := usuario{"Pedro", 24, endereco{"Rua Tal", 31}}

	fmt.Println(user, user2)
}
