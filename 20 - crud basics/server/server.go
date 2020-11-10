package server

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type user struct {
	ID uint32 `json:"id"`
	Name string `json:"nome"`
	Email string `json:"email"`
}

// CreateUser creates a new user within database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo de requisição"))
		return
	}

	var user user

	if err = json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar com banco de dados"))
		return
	}
	defer db.Close()

	// PREPARE STATEMENT - usado para evitar SQLInjection
	statement, err := db.Prepare("insert into users (nome, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insertion, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	insertedID, err := insertion.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter o ID inserido"))
		return
	}

	//STATUS CODE
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", insertedID)))




}

// SearchUsers returns all users
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	// SELECT * FROM USUARIOS
	lines, err := db.Query("select * from users")
	if err != nil {
		w.Write([]byte("Erro ao buscar os usuários!"))
		return
	}
	defer lines.Close()

	var users []user
	for lines.Next() {
		var user user
		if err := lines.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON!"))
		return
	}

}

// SearchUser return an user by ID
func SearchUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	line, err := db.Query("select * from users where id = ?", ID)
	if err != nil {
		w.Write([]byte("Erro ao buscar o usuário"))
		return
	}

	var user user
	if line.Next() {
		if err := line.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Erro ao converter o usuário para JSON"))
		return
	}

}

// UpdateUser update a user within database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"],10,32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição o parâmetro para inteiro"))
		return
	}

	var user user
	if err := json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update users set nome = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		w.Write([]byte("Erro ao atualizar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// DeleteUser deletes a user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"],10,32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from users where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Erro ao deletar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}