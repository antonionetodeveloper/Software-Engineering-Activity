package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	entities "student-information-system/internal/models/operations"
	"student-information-system/internal/models/operations/admin"
)

// Create handles requests to create a admin.
func Create(writer http.ResponseWriter, request *http.Request) {
	var user entities.Admin
	var response map[string]any

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest,
		)
		log.Println(err)
		return
	}

	id, err := admin.Create(user)
	if err != nil {
		response = map[string]any{
			"Success": false,
			"Message": fmt.Sprintf("An error ocurred while trying to instert user: %v", err),
		}
	} else {
		response = map[string]any{
			"Success": true,
			"Message": fmt.Sprintf("User successfuly added! User id: %d", id),
		}
		writer.WriteHeader(http.StatusCreated)
	}

	writer.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(response)
}
