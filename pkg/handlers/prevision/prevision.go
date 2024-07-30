package prevision

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/andersonnfreire/api-climate/pkg/apis"
	"github.com/andersonnfreire/api-climate/pkg/config"
	"github.com/go-playground/validator"
)

var validate = validator.New()

func GetPrevisionHandler(writerResponse http.ResponseWriter, request *http.Request, cfg config.Config) (*WeatherForecastsResponse, error) {

	params := ParamsGetPrevision{
		Token:    request.URL.Query().Get("token"),
		Cidade:   request.URL.Query().Get("cidade"),
		Language: request.URL.Query().Get("lang"),
	}

	if err := validate.Struct(params); err != nil {
		return nil, err
	}

	options := apis.HTTPRequestOptions{
		Method: http.MethodGet,
		URL:    "http://api.openweathermap.org/data/2.5/weather",
		QueryParams: map[string]string{
			"appid": params.Token,
		},
		Timeout: 5 * time.Minute,
	}

	if params.Cidade != "" {
		options.QueryParams["q"] = params.Cidade
	}

	if params.Language != "" {
		options.QueryParams["lang"] = params.Language
	}

	response, err := apis.SendHTTPRequest(options)
	if err != nil {
		return nil, err
	}

	var weatherData *WeatherForecastsResponse
	if err := json.Unmarshal(response.Body, &weatherData); err != nil {
		return nil, fmt.Errorf("erro ao decodificar dados da API: %s", err.Error())
	}

	return weatherData, nil
}
