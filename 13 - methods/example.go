package main

import "fmt"

// funções são coisas soltas, vc chama qd bem quiser, estão associadas ao pacote
// methods sempre estão associados a alguma coisa (struct, interface etc)
type user struct {
	name string
	age  uint8
}

func (u user) save() {
	// fmt.Println("Dentro do método salvar")
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.name)
}

func (u user) maiorDeIdade() bool { // dessa forma recebe os parâmetros como copia, qq alteração não refletirá na instância
	return u.age >= 18
}

func (u *user) fazerAniversario() { // qd quer alterar os parametros da struct tem q usar como ponteiro para refletir na instância
	u.age++ // não precisa desreferenciar qd usa dessa forma
}

func main() {
	user1 := user{"Lucas", 39}
	fmt.Println(user1)
	user1.save()

	user2 := user{"Daivide", 15}
	user2.save()

	fmt.Println(user1.maiorDeIdade())
	fmt.Println(user2.maiorDeIdade())

	user2.fazerAniversario()
	fmt.Println(user2.age)
}
