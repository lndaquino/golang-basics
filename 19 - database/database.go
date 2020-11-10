package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // importo drive aqui mas quem vai usar o drive é o pacote sql (import implicita)
)

func main() {
	// stringConexao := "usuario:senha@/banco?charset=utf8&parseTime=True&loc=Local"
	stringConexao := "golang:golang@tcp(192.168.99.100:3306)/usuarios?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão com DB aberta!")

	lines, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer lines.Close()

	fmt.Println(lines)
}
