package admin

import (
	"student-information-system/package/database"
)

func DeleteByID(id int64) (int64, error) {
	connection, err := database.OpenConnection()
	if err != nil {
		return 0, nil
	}
	defer connection.Close()

	SQL := `DELETE FROM admin WHERE id=$1`
	response, err := connection.Exec(SQL, id)
	if err != nil {
		return 0, nil
	}

	return response.RowsAffected()
}
