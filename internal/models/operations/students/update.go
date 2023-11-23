package students

import (
	"log"
	entities "student-information-system/internal/models/operations"
	"student-information-system/package/database"
)

func UpdateByID(id int64, user entities.Student) (int64, error) {
	connection, err := database.OpenConnection()
	if err != nil {
		return 0, nil
	}
	defer connection.Close()

	SQL := `UPDATE student SET name=$2, age=$3, cpf=$4, phone=$5, gender=$6 WHERE id=$1`
	response, err := connection.Exec(SQL, id, user.Name, user.Age, user.CPF, user.Phone, user.Gender)
	if err != nil {
		log.Println(err)
		return 0, nil
	}

	return response.RowsAffected()
}
