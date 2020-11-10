package main

import (
	"fmt"
	"tests-introduction/addresses"
)

func main() {
	addressType := addresses.Type("Travessa dos Imigrantes")
	fmt.Println(addressType)
}
