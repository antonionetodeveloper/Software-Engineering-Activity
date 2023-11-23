package admin

import (
	entities "student-information-system/internal/models/operations"
	"student-information-system/package/database"
)

func UpdateByID(id int64, user entities.Admin) (int64, error) {
	connection, err := database.OpenConnection()
	if err != nil {
		return 0, nil
	}
	defer connection.Close()

	SQL := `UPDATE admin SET name=$2, role=$3 WHERE id=$1`
	response, err := connection.Exec(SQL, id, user.Name, user.Role)
	if err != nil {
		return 0, nil
	}

	return response.RowsAffected()
}
