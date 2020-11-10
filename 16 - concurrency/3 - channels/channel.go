package main

// channel - é um canal de comunicação que pode enviar e receber dados, e usa esse envio para sincronizar as goroutines

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) // cria un canal que só pode enviar e receber dados no formato string
	go write("Olá Mundo", canal)

	fmt.Println("Depois da função write começar a ser executada")

	// for {
	// 	message, aberto := <-canal // espera q o canal receba um valor para continuar - deadlock ocorre qd vc não tem nenhum lugar enviando dado para seu canal (canal aberto), mas seu canal continua esperando um dado
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(message)
	// }

	for message := range canal {
		fmt.Println(message)
	}

	fmt.Println("Fim do programa")
}

func write(text string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- text
		time.Sleep(time.Second)
	}
	close(canal)
}
