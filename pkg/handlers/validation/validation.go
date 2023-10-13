package validation

import (
	"errors"
	"net/http"
)

func ValidateQueryParam(r *http.Request) error {
	apiKey := r.URL.Query().Get("token")
	city := r.URL.Query().Get("cidade")

	if city == "" {
		return errors.New("o parâmetro 'cidade' é obrigatório")
	}

	if apiKey == "" {
		return errors.New("o parâmetro 'token' é obrigatório")
	}

	return nil
}
