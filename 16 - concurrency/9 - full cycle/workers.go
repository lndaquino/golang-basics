package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	concurrency := 50
	input := make(chan int)
	done := make(chan []byte)

	go func() {
		i := 0
		for {
			input <- i
			i++
		}
	}()

	for x := 0; x < concurrency; x++ {
		go ProcessWorker(input, x)
	}
	<-done
}

// ProcessWorker blablabla
func ProcessWorker(input chan int, worker int) {
	for x := range input {
		t := time.Duration(rand.Intn(4) * int(time.Second))
		time.Sleep(t)
		fmt.Println("Worker ", worker, ": ", int(x))
	}
}
