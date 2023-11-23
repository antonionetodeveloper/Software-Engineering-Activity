package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"student-information-system/internal/models/operations/admin"
)

// GetByID handles requests to retrieve a student by ID.
func GetByID(writer http.ResponseWriter, request *http.Request) {
	response := make(map[string]interface{})

	idStringify := chi.URLParam(request, "id")
	id, err := strconv.Atoi(idStringify)
	if err != nil {
		handleError(writer, http.StatusBadRequest, err, "Error while parsing id", response)
		return
	}

	user, err := admin.GetByID(int64(id))
	if err != nil {
		handleError(writer, http.StatusNotFound, err, "Error while searching user", response)
		return
	}

	writer.WriteHeader(http.StatusOK)
	response["success"] = true
	response["user"] = user

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(response)
}

func handleError(writer http.ResponseWriter, status int, err error, logMessage string, response map[string]interface{}) {
	log.Println("Local: internal/handlers/read.go, handleError")
	log.Printf("%s: %v", logMessage, err)

	writer.WriteHeader(status)
	response["success"] = false

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(response)
}
