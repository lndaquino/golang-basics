package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays and Slices")
	// arrays tem tamanho fixo e é pouco usado
	// slice é baseado no array mas não tem tamanho fixo
	var array1 [5]string // se não declarar o tamanho o array vira um slice
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	// array2[5] = "Posição 6" // não deixa fazer pois o tamanho do array é fixo

	array3 := [...]int{1, 2, 3, 4, 5} // [...] infere o tamanho do array pela quantidade de valores declarados
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16} //slice não especificamos o tamanho fixo
	fmt.Println(slice)
	slice = append(slice, 17)
	fmt.Println(slice)

	slice2 := array2[1:3] // [inclusivo:exclusivo] ao pegar os valores - fica apontando para o array original (se altera array reflete no slice)
	slice2 = append(slice2, "Position")
	array2[1] = "Nova posição"
	fmt.Println(slice2)

	fmt.Println(reflect.TypeOf(slice)) // mostra o tipo de variável
	fmt.Println(reflect.TypeOf(array3))

	fmt.Println("-------")
	// Arrays internos
	slice3 := make([]float32, 10, 11) // cria um array com 11 posições e retorna um slice com as 10 primeiras posições do array
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity

	slice3 = append(slice3, 5) // atinge a capacidade máxima do slice
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity

	slice3 = append(slice3, 10) // quando vê que o seu array interno vai estourar o Go cria um novo array com o dobro da capacidade
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // length
	fmt.Println(cap(slice4)) // capacity

	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // length
	fmt.Println(cap(slice4)) // capacity
}
