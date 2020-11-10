package main

import (
	"crud/server"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// CRUD - CREATE (POST), READ (GET), UPDATE (PUT), DELETE
	router := mux.NewRouter()
	router.HandleFunc("/user", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/user", server.SearchUsers).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", server.SearchUser).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", server.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/user/{id}", server.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
