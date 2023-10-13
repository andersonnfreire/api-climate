package apis

import (
	"errors"
	"io"
	"net/http"
	"time"
)

type HTTPRequestOptions struct {
	Method      string
	URL         string
	Body        io.Reader
	Headers     map[string]string
	QueryParams map[string]string
	Timeout     time.Duration
	AuthToken   string
}

type HTTPResponse struct {
	StatusCode int
	Body       []byte
}

func SendHTTPRequest(options HTTPRequestOptions) (HTTPResponse, error) {

	if len(options.QueryParams) > 0 {
		// Adicione os parâmetros da query à URL.
		url := options.URL + "?"
		for key, value := range options.QueryParams {
			url += key + "=" + value + "&"
		}
		options.URL = url[:len(url)-1] // Remova o último "&" adicionado.
	}

	req, err := http.NewRequest(options.Method, options.URL, options.Body)
	if err != nil {
		return HTTPResponse{}, errors.New("Erro ao criar a solicitação: " + err.Error())
	}

	for key, value := range options.Headers {
		req.Header.Set(key, value)
	}

	if options.AuthToken != "" {
		req.Header.Set("Authorization", "Bearer "+options.AuthToken)
	}

	client := &http.Client{
		Timeout: options.Timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return HTTPResponse{
			StatusCode: http.StatusInternalServerError,
		}, errors.New("erro ao executar a solicitação: " + err.Error())
	}

	if resp == nil {
		return HTTPResponse{
			StatusCode: http.StatusInternalServerError,
		}, errors.New("nenhum retorno foi encontrado ao executar a solicitação")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return HTTPResponse{}, errors.New("erro ao ler o corpo da resposta: " + err.Error())
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return HTTPResponse{
			StatusCode: resp.StatusCode,
			Body:       body,
		}, errors.New("a solicitação falhou com o código de status " + resp.Status)
	}

	return HTTPResponse{
		StatusCode: resp.StatusCode,
		Body:       body,
	}, nil
}
