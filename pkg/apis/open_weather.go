package apis

import (
	"net/http"
	"time"

	"github.com/andersonnfreire/api-climate/pkg/handlers/validation"
)

// GetWeatherData faz uma requisição à API de previsão do tempo.
func GetWeatherData(request *http.Request, params validation.ParamsGetPrevision) (HTTPResponse, error) {
	options := HTTPRequestOptions{
		Method: http.MethodGet,
		URL:    "http://api.openweathermap.org/data/2.5/weather",
		QueryParams: map[string]string{
			"q":     params.Cidade,
			"appid": params.Token,
		},
		Timeout: 5 * time.Minute,
	}

	response, err := SendHTTPRequest(options)
	if err != nil {
		return HTTPResponse{}, err
	}

	return response, nil
}
