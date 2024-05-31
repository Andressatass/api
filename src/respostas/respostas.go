package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorta uma resposta em json para a requisição
func JSON(
	w http.ResponseWriter,
	statusCode int,
	dados interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		if err := json.NewEncoder(w).Encode(dados); err != nil {
			log.Fatal(err)
		}
	}
}

// Erro retorna um erro em formato JSON
func Erro(
	w http.ResponseWriter,
	statusCode int,
	erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
