package main

import (
	"fmt"
	"math"
)

// interfaces são mt usadas qd precisamos ter uma certa flexibilidade com os tipos
// struct tem campos
// interface tem assinatura de metodos
// a implementação do metodo na struct é implicita
type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

// tipos genéricos
func generica(inter interface{}) {
	fmt.Println(inter)
}

func main() {
	r := retangulo{10, 15}
	c := circulo{5}

	escreverArea(r)
	escreverArea(c)

	// passando parâmetro para interface genérica que aceitaria qq coisa - ex: Println recebe esse tipo de parêmetro
	generica("string")
	generica(1)
	generica(true)
	generica(r)

	mapa := map[interface{}]interface{}{
		1:        true,
		"String": float64(3.141592),
		true:     false,
		uint(8):  float32(95405.2),
	}

	fmt.Println(mapa)
}
