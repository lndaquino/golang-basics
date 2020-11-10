package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // informa a qtd de goroutines que precisará colocar na fila

	go func() { // cria primeira goroutine para rodar de forma concorrente
		write("Hello World!")
		waitGroup.Done() // depois q função terminar chama done para dizer q aquela goroutine terminou (-1 no contador)
	}()

	go func() {
		write("Programing in Go!")
		waitGroup.Done()
	}()

	waitGroup.Wait() // informa main pra esperar contagem chegar em zero
}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
