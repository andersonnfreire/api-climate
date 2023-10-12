package prevision

import "net/http"

func GetPrevisionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Olá, Teste!"))
}
