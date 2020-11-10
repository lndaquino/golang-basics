// Unit test
package addresses

import "testing"

type testScenario struct {
	param  string
	result string
}

func TestType(t *testing.T) { //função de teste precisa começar com TestQqCoisa
	// testAddress := "Rua Paulista"

	// addressType := "Rua"

	// responseAddressType := Type(testAddress)

	// if responseAddressType != addressType {
	// 	t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", addressType, responseAddressType)
	// }

	// usando cenários de testes
	t.Parallel()

	scenarios := []testScenario{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo inválido"},
		{"Estrada das Almas", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"rua sem nome", "Rua"},
		{"AVENIDA ABC", "Avenida"},
		{"", "Tipo inválido"},
	}

	for _, scenario := range scenarios {
		response := Type(scenario.param)
		if response != scenario.result {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", scenario.result, response)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}
