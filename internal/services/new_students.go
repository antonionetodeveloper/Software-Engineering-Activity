package services

import (
	"encoding/json"
	"net/http"
	entities "student-information-system/internal/models/operations"
	"student-information-system/package/database"
)

// ShowNewStudents show the lasts signed up students.
func ShowNewStudents(writer http.ResponseWriter, request *http.Request) {
	var students []entities.Student

	connection, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	SQL := `SELECT * FROM student ORDER BY ID DESC`
	rows, err := connection.Query(SQL)
	if err != nil {
		return
	}

	for rows.Next() {
		var user entities.Student

		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.CPF, &user.Phone, &user.Gender)
		students = append(students, user)
	}

	response := map[string]any{
		"Success":  true,
		"students": students,
	}

	writer.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(response)
}
