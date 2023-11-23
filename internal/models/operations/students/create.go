package students

import (
	entities "student-information-system/internal/models/operations"
	"student-information-system/package/database"
)

func Create(user entities.Student) (id int64, err error) {
	connection, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	SQL := `INSERT INTO student (name, age, cpf, phone, gender) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = connection.QueryRow(SQL, user.Name, user.Age, user.Gender, user.CPF, user.Phone).Scan(&id)

	return
}
