package main

import (
	"errors"
	"fmt"
)

func main() {
	// int, int8, int16, int32, int64
	// int atribui o valor conforme a arquitetura do computador (int64 se estiver em computador 64bits ou int32)
	var numero int = 10000000000000
	fmt.Println(numero)

	// uint - unsigned int - não aceita valores negativos - tb tem uint8, uint6, uint32, uint64
	var num2 uint = 1000
	fmt.Println(num2)

	// alias
	// int32 = rune
	var num3 rune = 1000

	// int8 = byte
	var num4 byte = 128

	// numeros reaus
	var num5 float32 = 123.45
	var num6 float64 = 123456.789

	num7 := 1123454.4885
	fmt.Println(num3, num4, num5, num6, num7)

	var str string = "Texto"
	fmt.Println(str)

	// aspas simples indicam que é um caracter q é convertido para o codigo ASCII, aspas duplas é string
	caracter := 'l'
	fmt.Println(caracter)

	// valor zero = todo tipo de dado tem seu valor zero caso a variável não seja inicializada (string nula, int 0, error nil, bool false)
	var texto int
	fmt.Println(texto)

	var flag bool
	fmt.Println(flag)

	var err error = errors.New("Erro interno")
	fmt.Println(err)
}
