package main

import (
	config "student-information-system/configs"
	"student-information-system/internal/api"
)

func main() {
	go func() {
		api.Router()
	}()
	config.CheckStopServer()
}
