package main

import "fmt"

func main() {
	var var1 string = "Lucas"
	var2 := "Nóbrega"
	var (
		var3 string = "Dantas"
		var4 string = "de"
	)
	var5, var6 := "Aquino", "OTFSSPX"

	fmt.Println(var1 + " " + var2 + " " + var3 + " " + var4 + " " + var5 + " " + var6)

	const const1 string = "constante"
	fmt.Println(const1)

	// invertendo valor entre variáveis, pode fazer direto como em Python
	var5, var6 = var6, var5
	fmt.Println(var1 + " " + var2 + " " + var3 + " " + var4 + " " + var5 + " " + var6)
}
