package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplex(write("Olá Mundo!"), write("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplex(canal1, canal2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case msg1 := <-canal1:
				canalDeSaida <- msg1
			case msg2 := <-canal2:
				canalDeSaida <- msg2
			}
		}
	}()

	return canalDeSaida
}

func write(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
