package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go só tem for, não tem while, do while")
	i := 0
	for i < 10 {
		fmt.Println("Incrementando i")
		// time.Sleep(time.Second)
		i++
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
		// time.Sleep(time.Second)
	}
	// fmt.Println(j) // não dá pra acessar fora do escopo do for

	nomes := [3]string{"Lucas", "Juliana", "Maria"}

	for indice, nome := range nomes { // pra ignorar o indice usar _ no lugar do indice
		fmt.Println(indice, nome)
		// time.Sleep(time.Second)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	user := map[string]string{ // range só dá pra usar em map, não em struct
		"nome":  "Lucas",
		"idade": "11",
	}

	for chave, valor := range user {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}
