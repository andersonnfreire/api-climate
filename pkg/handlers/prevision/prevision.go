package prevision

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/andersonnfreire/api-climate/pkg/apis"
	"github.com/andersonnfreire/api-climate/pkg/config"
	"github.com/andersonnfreire/api-climate/pkg/handlers/validation"
)

func GetPrevisionHandler(writerResponse http.ResponseWriter, request *http.Request, cfg config.Config) {

	if err := validation.ValidateQueryParam(request); err != nil {
		errorMsg := map[string]string{
			"error": err.Error(),
		}
		response, err := json.Marshal(errorMsg)
		if err != nil {
			http.Error(writerResponse, "Erro interno do servidor", http.StatusInternalServerError)
			return
		}
		writerResponse.Header().Set("Content-Type", "application/json")
		writerResponse.WriteHeader(http.StatusUnprocessableEntity)
		writerResponse.Write(response)
		return
	}

	response, err := apis.SendHTTPRequest(apis.HTTPRequestOptions{
		Method: "GET",
		URL:    "http://api.openweathermap.org/data/2.5/weather",
		Headers: map[string]string{
			"q":     request.URL.Query().Get("cidade"),
			"appid": request.URL.Query().Get("token"),
		},
		Timeout: 5 * time.Minute,
	})

	writerResponse.Header().Set("Content-Type", "application/json")
	writerResponse.WriteHeader(response.StatusCode)

	if err != nil {
		errorMsg := map[string]string{
			"error": err.Error(),
		}
		responseMsg, err := json.Marshal(errorMsg)
		if err != nil {
			http.Error(writerResponse, "Erro interno do servidor", http.StatusInternalServerError)
			return
		}
		writerResponse.Write(responseMsg)
	}
	writerResponse.Write([]byte(response.Body))
}
