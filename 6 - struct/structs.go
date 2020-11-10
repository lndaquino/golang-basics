package main

import "fmt"

// struct é uma coleção de campos com variáveis e tipos
type user struct {
	nome    string
	idade   uint8
	address endereco
}

type endereco struct {
	rua    string
	numero uint8
}

func main() {
	fmt.Println("Structs")

	var user1 user
	user1.nome = "Lucas"
	user1.idade = 18
	fmt.Println(user1)

	address1 := endereco{"Rua dos bobos", 0}

	user2 := user{"Davi", 21, address1}
	fmt.Println(user2)

	user3 := user{nome: "José"}
	fmt.Println(user3)
}
