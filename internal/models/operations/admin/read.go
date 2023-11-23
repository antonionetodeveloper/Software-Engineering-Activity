package admin

import (
	entities "student-information-system/internal/models/operations"
	"student-information-system/package/database"
)

func GetByID(id int64) (user entities.Admin, err error) {
	connection, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	SQL := `SELECT * FROM admin WHERE id=$1;`
	row := connection.QueryRow(SQL, id)
	err = row.Scan(&user.ID, &user.Name, &user.Role)

	return
}
