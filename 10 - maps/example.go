package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[int]string{
		1: "Pedro",
		2: "Silva",
	}

	fmt.Println(usuario[1])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Lucas",
			"ultimo":   "Aquino",
		},
		"curso": {
			"nome":      "Engenharia",
			"faculdade": "USP",
		},
	}

	fmt.Println(usuario2["nome"]["primeiro"])
	fmt.Println(usuario2)
	delete(usuario2, "nome") // apaga a chave do map
	fmt.Println(usuario2)

	// incluindo uma nova chave no map
	usuario2["contato"] = map[string]string{
		"telefone": "1111111",
	}
	fmt.Println(usuario2)

}
