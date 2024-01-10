package server

import (
	"fmt"
	"net/http"

	"github.com/andersonnfreire/api-climate/pkg/apis"
	"github.com/andersonnfreire/api-climate/pkg/config"
	"github.com/andersonnfreire/api-climate/pkg/handlers/pdf"
	"github.com/andersonnfreire/api-climate/pkg/handlers/prevision"
)

// APIHandler manipula a requisição à API e gera o PDF.
func APIHandler(writeResponse http.ResponseWriter, request *http.Request, cfg config.Config) {
	// Faça a requisição à API de previsão do tempo.
	weatherData, err := prevision.GetPrevisionHandler(writeResponse, request, cfg)
	if err != nil {
		apis.ResponseWithError(writeResponse, http.StatusInternalServerError, fmt.Sprintf("Erro ao obter dados da API: %s", err.Error()))
		return
	}

	// Gere o PDF com base nos dados da previsão do tempo.
	if err := pdf.GeneratePDF(weatherData); err != nil {
		apis.ResponseWithError(writeResponse, http.StatusInternalServerError, fmt.Sprintf("Erro ao gerar o PDF: %s", err.Error()))
		return
	}
	apis.Response(writeResponse, http.StatusOK, []byte("PDF gerado com sucesso!"))
}
