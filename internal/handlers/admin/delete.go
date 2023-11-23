package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"student-information-system/internal/models/operations/admin"
)

// DeleteByID handles requests to retrieve a admin by ID.
func DeleteByID(writer http.ResponseWriter, request *http.Request) {
	response := make(map[string]interface{})

	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("Error while parsing id: %v", err)
		http.Error(writer,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	_, err = admin.DeleteByID(int64(id))
	if err != nil {
		log.Printf("Error while deleting user: %v", err)
		http.Error(writer,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
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
