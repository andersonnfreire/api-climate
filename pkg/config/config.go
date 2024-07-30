package config

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host           string
	Port           string
	UrlOpenWeather string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		// log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	url := os.Getenv("URL_OPEN_WEATHER")
	if url == "" {
		url = "http://api.openweathermap.org/data/2.5/weather"
	}

	return &Config{
		Host:           host,
		Port:           port,
		UrlOpenWeather: url,
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
