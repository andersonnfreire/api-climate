package routes

import (
	"net/http"

	"github.com/andersonnfreire/api-climate/pkg/handlers/prevision"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", prevision.GetPrevisionHandler)
}
