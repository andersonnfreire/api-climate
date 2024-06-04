package server

import (
	"fmt"
	"net/http"

	"github.com/andersonnfreire/api-climate/pkg/apis"
	"github.com/andersonnfreire/api-climate/pkg/config"
	"github.com/andersonnfreire/api-climate/pkg/handlers/prevision"
	"github.com/andersonnfreire/api-climate/pkg/handlers/template"
)

// ServerHandler manipula a requisição à API e gera o HTML.
func ServerHandler(writeResponse http.ResponseWriter, request *http.Request, cfg config.Config) {
	// Faça a requisição à API de previsão do tempo.
	weatherData, err := prevision.GetPrevisionHandler(writeResponse, request, cfg)
	if err != nil {
		apis.ResponseWithError(writeResponse, http.StatusInternalServerError, fmt.Sprintf("Erro ao obter dados da API: %s", err.Error()))
		return
	}

	// Gere o HTML com base nos dados da previsão do tempo.
	_, err = template.RenderTemplateHTML(weatherData)
	if err != nil {
		apis.ResponseWithError(writeResponse, http.StatusInternalServerError, fmt.Sprintf("Erro ao gerar o HTML: %s", err.Error()))
		return
	}

	apis.Response(writeResponse, http.StatusOK, []byte("HTML gerado com sucesso!"))
}
