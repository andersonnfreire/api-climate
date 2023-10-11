package routes

import (
	"net/http"

	"github.com/andersonnfreire/api-climate/pkg/controllers"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", controllers.IndexHandler)
	mux.HandleFunc("/about", controllers.AboutHandler)
}
