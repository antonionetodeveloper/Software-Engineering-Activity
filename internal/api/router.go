package api

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	config "student-information-system/configs"
	admin "student-information-system/internal/handlers/admin"
	students "student-information-system/internal/handlers/students"
	"student-information-system/internal/services"
)

// Router configure and return a new HTTP router.
func Router() {
	err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Printf("Server running on port: %v", config.GetServerPort())

	router := chi.NewRouter()

	// Admin
	router.Get("/admin/{id}", admin.GetByID)
	router.Post("/admin/create", admin.Create)
	router.Put("/admin/edit/{id}", admin.UpdateUser)
	router.Delete("/admin/delete/{id}", admin.DeleteByID)

	// Students
	router.Get("/student/{id}", students.GetByID)
	router.Post("/student/create", students.Create)
	router.Put("/student/edit/{id}", students.UpdateByID)
	router.Delete("/student/delete/{id}", students.DeleteByID)

	// Services
	router.Get("/services/show-new-students", services.ShowNewStudents)

	_ = http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), router)
}
