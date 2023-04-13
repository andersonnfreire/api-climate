package config

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Erro ao carregar o arquivo .env: %v", err)
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Host: host,
		Port: port,
	}, nil
}

func (c *Config) ServerAddress() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func (c *Config) Server() *http.Server {
	return &http.Server{
		Addr: c.ServerAddress(),
	}
}
