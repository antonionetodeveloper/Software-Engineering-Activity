package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	entities "student-information-system/internal/models/operations"
	"student-information-system/internal/models/operations/admin"
)

// UpdateUser handles requests to update a admin by ID.
func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	var user entities.Admin

	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("Error while parsing id: %v", err)
		http.Error(writer,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}

	_, err = admin.UpdateByID(int64(id), user)
	if err != nil {
		log.Printf("Error while updating user: %v", err)
		http.Error(writer,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}

	response := map[string]any{
		"Success": true,
		"Message": fmt.Sprintf("User successfuly updated! User id: %d", id),
	}

	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
