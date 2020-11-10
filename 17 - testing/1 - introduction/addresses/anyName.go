package addresses

import "strings"

//Type returns if the address has a valid a type and return its type
func Type(address string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}

	addressLowerCase := strings.ToLower(address)
	firstWord := strings.Split(addressLowerCase, " ")[0]

	validAddress := false
	for _, tipo := range validTypes {
		if tipo == firstWord {
			validAddress = true
		}
	}

	if validAddress {
		return strings.Title(firstWord)
	}

	return "Tipo inválido"

}
