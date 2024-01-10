package apis

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func Response(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(body)
}

// RespondWithError envia uma resposta de erro JSON ao cliente.
func ResponseWithError(w http.ResponseWriter, statusCode int, errMsg string) {
	errResponse := ErrorResponse{Error: errMsg}
	body, _ := json.Marshal(errResponse)
	Response(w, statusCode, body)
}
