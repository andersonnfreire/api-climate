package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Registra um handler para a rota "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Bem-vindo ao meu servidor local!")
	})

	// Inicia o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}
