package main

import (
	"fmt"
	"time"
)

func fibRecursive(posicao int) int { // não se sugere funções recursivas para valores grandes pois pode dar estouro de pilha
	if posicao <= 1 {
		return posicao
	}
	return fibRecursive(posicao-1) + fibRecursive(posicao-2)
}

func worker(tarefas <-chan int, resultados chan<- int) { // especifica q canal só recebe dados (<-chan) ou só envia dados (chan<-), ou seja, deixam de ser canais bidirecionais
	for numero := range tarefas {
		resultados <- fibRecursive(numero)
	}
}

func main() {
	start := time.Now()

	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	position := 45

	for i := 1; i <= position; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 1; i <= position; i++ {
		resultado := <-resultados
		fmt.Printf("%d --> %d\n", i, resultado)
	}

	end := time.Now()
	fmt.Println(end.Sub(start))
}
