package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"student-information-system/internal/models/operations/students"
)

const fileRead = "Local: internal/handlers/read.go"

// GetByID handles requests to retrieve a student by ID.
func GetByID(writer http.ResponseWriter, request *http.Request) {
	response := make(map[string]interface{})

	idStringify := chi.URLParam(request, "id")
	id, err := strconv.Atoi(idStringify)
	if err != nil {
		handleError(
			writer, http.StatusBadRequest, err, fileRead, "Error while parsing id", response,
		)
		return
	}

	student, err := students.GetByID(int64(id))
	if err != nil {
		handleError(
			writer, http.StatusNotFound, err, fileRead, "Error while searching user", response,
		)
		return
	}

	writer.WriteHeader(http.StatusOK)
	response["success"] = true
	response["user"] = student

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(response)
}

func handleError(writer http.ResponseWriter, status int, err error, fileError string, logMessage string, response map[string]interface{}) {
	log.Println(fileError)
	log.Printf("%s: %v", logMessage, err)

	writer.WriteHeader(status)
	response["success"] = false

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(response)
}
