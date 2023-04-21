package controllers

import (
	"encoding/json"
	"net/http"

	models "github.com/andersonnfreire/api-climate/pkg/models/user"
	utils "github.com/andersonnfreire/api-climate/pkg/utils"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Olá, mundo!"))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	users := []models.User{
		{Name: utils.ToUpper("Usuário 1"), Email: "usuario1@exemplo.com"},
		{Name: utils.ToUpper("Usuário 2"), Email: "usuario2@exemplo.com"},
		{Name: utils.ToUpper("Usuário 3"), Email: "usuario3@exemplo.com"},
	}

	// Serializa a lista de usuários como JSON
	data, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao serializar usuários como JSON"))
		return
	}

	// Escreve a lista de usuários como JSON na resposta HTTP
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
