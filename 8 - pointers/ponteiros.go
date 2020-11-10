package main

import "fmt"

func inverterSinal(numero int) int { // passando parametro por cópia
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) { // passando paramentro por referência, qualquer alteração altera a variável referenciada
	*numero = *numero * -1
}

func main() {
	fmt.Println("Ponteiro é uma referência de memória")

	var v1 int = 10
	var v2 int = v1 // copia o valor da variável para a nova variável
	// daqui pra baixo v1 e v2 não tem mais nenhuma ligação entre si
	fmt.Println(v1, v2)
	v1++
	fmt.Println(v1, v2)

	var v3 int
	var ponteiro *int // valor zero é nil - guarda o endereço de memória de um inteiro
	fmt.Println(v3, ponteiro)

	v3 = 100
	ponteiro = &v3
	fmt.Println(v3, ponteiro)  // ponteiro mostra o endereço da memória q está apontando
	fmt.Println(v3, *ponteiro) // desreferenciação, ponteiro vai mostrar o valor q está salvo no endereço apontado

	v3 = 150
	fmt.Println(v3, ponteiro)  // ponteiro mostra o endereço da memória q está apontando
	fmt.Println(v3, *ponteiro) // desreferenciação, ponteiro vai mostrar o valor q está salvo no endereço apontado

	numero := 20
	numeroInvertido := inverterSinal(numero) // faz copia do valor numero em uma nova variável
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero) // passa o endenreço de memória
	fmt.Println(novoNumero)
}
