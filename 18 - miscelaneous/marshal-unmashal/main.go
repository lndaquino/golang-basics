package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// struct com campos mapeados para json
type dog struct {
	Name string `json:"name"`
	Raca string `json:"-"` // não aparece na conversão pra JSON
	Age  uint8  `json:"chaveQualquer"`
}

func main() {
	// convertendo struct pra JSON
	dog1 := dog{"Rex", "Dálmata", 3}
	fmt.Println(dog1)

	dog1JSON, err := json.Marshal(dog1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dog1JSON)                  // devolve um slice de bytes []uint8
	fmt.Println(bytes.NewBuffer(dog1JSON)) // devolve um slice de bytes []uint8

	// convertendo map pra JSON
	dog2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	dog2JSON, err := json.Marshal(dog2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dog2JSON)                  // devolve um slice de bytes []uint8
	fmt.Println(bytes.NewBuffer(dog2JSON)) // devolve um slice de bytes []uint8

	// convertendo JSON para struct
	dog3JSON := `{"name":"Rex","raca":"Dálmata","chaveQualquer":3}`
	var dog3 dog
	if err := json.Unmarshal([]byte(dog3JSON), &dog3); err != nil {
		log.Fatal(err)
	}

	fmt.Println(dog3)

	// convertendo JSON para map
	dog4JSON := `{"nome":"Toby","raca":"Poodle"}`
	dog4 := make(map[string]string)

	if err := json.Unmarshal([]byte(dog4JSON), &dog4); err != nil {
		log.Fatal(err)
	}

	fmt.Println(dog4)

}
