package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calc(p1, p2 int8) (int8, int8) {
	soma := p1 + p2
	sub := p1 - p2

	return soma, sub
}

// função com retorno nomeado
func calc2(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

// função variática - pode receber vários parâmetros
func somaVariatica(numeros ...int) int { // pode receber de 0 a n ints
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

func escrever(texto string, numeros ...int) { // variavel variatica deve ser sempre o último parâmetro da função
	for _, num := range numeros {
		fmt.Println(texto, num)
	}
}

func fibRecursive(posicao uint) uint { // não se sugere funções recursivas para valores grandes pois pode dar estouro de pilha
	if posicao <= 1 {
		return posicao
	}
	return fibRecursive(posicao-1) + fibRecursive(posicao-2)
}

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoAprovado(n1, n2 float64) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!") // adia a execução até imediatamente antes do retorno da função
	fmt.Println("Verificando media do aluno")

	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func recuperarExecução() {
	fmt.Println("Tentativa de recuperar a execução")
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func mediaAluno(n1, n2 float64) bool {
	defer recuperarExecução()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!") // para de executar a função e entra em modo panico - vai chamar todas as funções que tem defer e depois mata o programa se não usar recover
}

// função Closure
func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

var n int

func init() { // pode ter uma por arquivo, não necessariamente por pacote, usada pra setup (inicialização de variável global etc)
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Executando a função main")
	fmt.Println(n)
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	res := f("Função 1")
	fmt.Println(res)

	// underline ignora retorno
	resulSoma, _ := calc(10, 5)
	fmt.Println(resulSoma)

	resulSum, resulSub := calc2(20, 10)
	fmt.Println(resulSum, resulSub)

	totalSoma := somaVariatica(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(totalSoma)

	escrever("Olá mundo", 10, 20, 30, 40)

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando o parâmetro") // declara e executa função anonima com o parâmetro

	fmt.Println(retorno)

	retorno2 := ("mais um chamado pra função anônima")
	fmt.Println(retorno2)

	posicao := uint(15)
	fmt.Println(fibRecursive(posicao))

	for i := uint(1); i < posicao; i++ {
		fmt.Println(fibRecursive(i))
	}

	defer funcao1() // defer = adiar --> adia a execução dessa linha até imediatamente antes do retorno/fim da execução da função
	funcao2()
	fmt.Println(alunoAprovado(1, 8))

	fmt.Println(mediaAluno(8, 6))
	fmt.Println("Pós mediaAluno")

	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()

}
