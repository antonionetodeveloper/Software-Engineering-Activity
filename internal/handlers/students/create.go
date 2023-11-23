package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	entities "student-information-system/internal/models/operations"
	"student-information-system/internal/models/operations/students"
	"student-information-system/package/utils"
)

const fileCreate = "Local: internal/handlers/create.go"

// Create handles requests to create a student.
func Create(writer http.ResponseWriter, request *http.Request) {
	var user entities.Student
	var response map[string]any

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		handleError(
			writer, http.StatusBadRequest, err, fileCreate, "Error while decoding json.", response,
		)
	}

	if !utils.CheckIsOfAge(user.Age) {
		handleError(
			writer, http.StatusBadRequest, err, fileCreate, "The user is not of age.", response,
		)
	}

	id, err := students.Create(user)
	if err != nil {
		handleError(
			writer, http.StatusBadRequest, err, fileCreate, "Could not create user.", response,
		)
	} else {
		response = map[string]any{
			"Message": fmt.Sprint("You successfully got in! (payment ticket)"),
			"UserId":  id,
			"Success": true,
		}
		writer.WriteHeader(http.StatusCreated)
	}

	writer.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(response)
}
