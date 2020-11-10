package main

import "fmt"

func main() {
	// aritméticos
	soma := 1 + 2 // não pode somar, subtrair, comparar etc tipos diferentes - não pode somar int16 + int32
	subtracao := 5 - 3
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// atribuição
	var variavel1 string = "string"
	var2 := "string"

	fmt.Print(variavel1, var2)

	// relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// operadores lógicos
	fmt.Println(true && true)
	fmt.Println(true || true)
	fmt.Println(!true)

	// operadores unários
	numero := 10
	numero++
	numero += 10

	numero--
	numero -= 20

	numero *= 3
	numero /= 2
	numero %= 3
	fmt.Println(numero)

	// operadores ternários não existem em Go, tem q usar o if/else
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
}
