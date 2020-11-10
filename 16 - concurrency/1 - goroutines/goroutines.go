package main

import (
	"fmt"
	"time"
)

func main() {
	// concorrência != paralelismo (executa ao mesmo tempo em nucleos diferentes)
	go write("Hello World!")   // executa essa linha mas não espera ele terminar para continuar executando o resto do programa
	write("Programing in Go!") // se colocar go aqui tb o programa acaba e não executa nenhum dos dois
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
