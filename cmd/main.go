package main

import (
	"log"
	"net/http"

	"github.com/andersonnfreire/api-climate/pkg/config"
	"github.com/andersonnfreire/api-climate/pkg/routes"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	server := cfg.Server()

	// Cria o ServeMux e adiciona as rotas
	mux := http.NewServeMux()
	routes.AddRoutes(mux, *cfg)

	// Define o handler do servidor como o ServeMux
	server.Handler = mux

	log.Printf("Servidor iniciado em http://%s", cfg.ServerAddress())

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
