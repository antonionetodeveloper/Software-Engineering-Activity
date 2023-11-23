package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	entities "student-information-system/internal/models/operations"
	"student-information-system/internal/models/operations/students"
)

const fileUpdate = "Local: internal/handlers/update.go"

// UpdateByID handles requests to update a student by ID.
func UpdateByID(writer http.ResponseWriter, request *http.Request) {
	var user entities.Student
	response := make(map[string]interface{})

	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		handleError(
			writer, http.StatusBadRequest, err, fileUpdate, "Error while parsing id.", response,
		)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		handleError(
			writer, http.StatusNotFound, err, fileUpdate, "Error while searching user", response,
		)
		return
	}

	_, err = students.UpdateByID(int64(id), user)
	if err != nil {
		handleError(
			writer, http.StatusInternalServerError, err, fileUpdate, "Error while updating user.", response,
		)
		return
	}

	response = map[string]any{
		"Success": true,
		"Message": fmt.Sprintf("User successfuly updated! User id: %d", id),
	}

	writer.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(response)
}
