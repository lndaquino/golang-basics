package forms

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	//TDD - Test Driven Development - escreve o teste primeiro, depois escreve o minimo de teste possível para fazer a função passar no teste, e vai incrementando os testes e função
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		expectedArea := float64(120)
		testedArea := ret.Area()

		if expectedArea != testedArea {
			t.Errorf("A área recebida %f é diferente da esperada %f", testedArea, expectedArea) // continua os testes com Error
			// t.Fatalf("A área recebida %f é diferente da esperada %f", testedArea, expectedArea) // para os testes se chegar no Fatal
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		radius := float64(10)
		circ := Circulo{radius}
		expectedArea := float64(math.Pi * math.Pow(radius, 2))
		testedArea := circ.Area()

		if expectedArea != testedArea {
			t.Errorf("A área recebida %f é diferente da esperada %f", testedArea, expectedArea) // continua os testes com Error
			// t.Fatalf("A área recebida %f é diferente da esperada %f", testedArea, expectedArea) // para os testes se chegar no Fatal
		}
	})
}
