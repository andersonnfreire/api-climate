package prevision

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andersonnfreire/api-climate/pkg/apis"
	"github.com/andersonnfreire/api-climate/pkg/config"
	"github.com/andersonnfreire/api-climate/pkg/handlers/validation"
)

func GetPrevisionHandler(writerResponse http.ResponseWriter, request *http.Request, cfg config.Config) (*WeatherForecastsResponse, error) {
	params, err := validation.ValidateQueryParamGetPrevision(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := apis.GetWeatherData(request, params)
	if err != nil {
		return nil, err
	}

	var weatherData *WeatherForecastsResponse
	if err := json.Unmarshal(apiResponse.Body, &weatherData); err != nil {
		return nil, fmt.Errorf("erro ao decodificar dados da API: %s", err.Error())
	}

	return weatherData, nil
}
