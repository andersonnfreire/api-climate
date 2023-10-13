package routes

import (
	"net/http"

	"github.com/andersonnfreire/api-climate/pkg/config"
	"github.com/andersonnfreire/api-climate/pkg/handlers/prevision"
)

func AddRoutes(mux *http.ServeMux, cfg config.Config) {
	mux.HandleFunc("/search/weather/forecast", func(write http.ResponseWriter, request *http.Request) {
		prevision.GetPrevisionHandler(write, request, cfg)
	})
}
