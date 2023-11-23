package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateTicket(writer http.ResponseWriter, request *http.Request) {
	response := map[string]any{
		"Success": true,
		"Message": fmt.Sprintf("Você ingressou com sucesso!"),
	}

	_ = json.NewEncoder(writer).Encode(response)
}
