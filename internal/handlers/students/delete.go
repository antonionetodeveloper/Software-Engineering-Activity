package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"student-information-system/internal/models/operations/students"
)

const fileDelete = "Local: internal/handlers/delete.go"

// DeleteByID handles requests to delete a student.
func DeleteByID(writer http.ResponseWriter, request *http.Request) {
	response := make(map[string]interface{})

	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		handleError(
			writer, http.StatusBadRequest, err, fileDelete, "Error while parsing id.", response,
		)
		return
	}

	_, err = students.DeleteByID(int64(id))
	if err != nil {
		handleError(
			writer, http.StatusInternalServerError, err, fileDelete, "Error while deleting user.", response,
		)
		return
	}

	response = map[string]any{
		"Success": true,
		"Message": fmt.Sprintf("User successfuly deleted!"),
	}

	writer.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(response)
}
