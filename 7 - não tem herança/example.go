package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa    // só passa o tipo, está pegando todos os campos de pessoa e jogando em estudante (como copiar colar)
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Go não tem herança nativamente")

	p1 := pessoa{"João", "Pedro", 20, 178}

	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1.nome, e1.curso)
	fmt.Println(e1.pessoa.nome, e1.curso)
}
