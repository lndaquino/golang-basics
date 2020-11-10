package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := -11

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menou ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 { // outroNumero é vista somente no escopo dentro do if - inicializa e usa
		fmt.Println("Número é maior do que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor do que -10")
	} else {
		fmt.Println("Entre 0 e -10")

	}
	// fmt.Println(outroNumero) // variável não é vista nesse nível

	// SWITCH
	dia := diaDaSemana(200)
	fmt.Println(dia)

	dia2 := diaDaSemana2(5)
	fmt.Println(dia2)

}

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string {
	var diadaSemana string

	switch {
	case numero == 1:
		diadaSemana = "Domingo"
		fallthrough // força a entrar no próximo case depois desse - tb não tem break, executou a condição sai do switch
	case numero > 1 && numero < 8:
		diadaSemana = "Dia válido"

	default:
		diadaSemana = "Número inválido"
	}

	return diadaSemana
}
