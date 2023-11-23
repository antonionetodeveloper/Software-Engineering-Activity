package admin

import (
	entities "student-information-system/internal/models/operations"
	"student-information-system/package/database"
)

func Create(user entities.Admin) (id int64, err error) {
	connection, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	SQL := `INSERT INTO admin (name, role) VALUES ($1, $2) RETURNING id`
	err = connection.QueryRow(SQL, user.Name, user.Role).Scan(&id)

	return
}
