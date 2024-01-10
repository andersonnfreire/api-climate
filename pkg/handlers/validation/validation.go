package validation

import (
	"net/http"

	"github.com/go-playground/validator"
)

var validate = validator.New()

type ParamsGetPrevision struct {
	Token  string `validate:"required"`
	Cidade string `validate:"required"`
}

func ValidateQueryParamGetPrevision(r *http.Request) (ParamsGetPrevision, error) {
	params := ParamsGetPrevision{
		Token:  r.URL.Query().Get("token"),
		Cidade: r.URL.Query().Get("cidade"),
	}
	if err := validate.Struct(params); err != nil {
		return ParamsGetPrevision{}, err
	}
	return params, nil
}
