package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	serverConfig "student-information-system/configs"
)

func OpenConnection() (*sql.DB, error) {
	config := serverConfig.GetDB()
	stringConnection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Pass, config.Database,
	)

	connection, err := sql.Open("postgres", stringConnection)
	if err != nil {
		panic(err)
	}

	err = connection.Ping()

	return connection, err
}
