package prevision

import (
	"net/http"

	"github.com/andersonnfreire/api-climate/pkg/apis"
	"github.com/andersonnfreire/api-climate/pkg/config"
	"github.com/andersonnfreire/api-climate/pkg/handlers/validation"
)

func GetPrevision(writerResponse http.ResponseWriter, request *http.Request, cfg config.Config) (apis.HTTPResponse, error) {
	params, err := validation.ValidateQueryParamGetPrevision(request)
	if err != nil {
		return apis.HTTPResponse{}, err
	}

	apiResponse, err := apis.GetWeatherData(request, params)
	if err != nil {
		return apiResponse, err
	}
	return apiResponse, nil
}
