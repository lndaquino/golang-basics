package main

import "fmt"

func main() {
	canal := make(chan string, 2) // canal com buffer, especifica a qtd máxima de mensagens antes do deadlock (capacidade máxima)

	canal <- "Olá Mundo\n"       // dá deadlock pois está enviando valor sem ninguém para receber, é operação bloqueante. Por isso se usa canais em funções diferentes
	canal <- "Programando em Go" // dá deadlock pois está enviando valor sem ninguém para receber, é operação bloqueante. Por isso se usa canais em funções diferentes
	mensagem := <-canal
	msg2 := <-canal
	// msg3 := <-canal

	fmt.Println(mensagem, msg2)
}
