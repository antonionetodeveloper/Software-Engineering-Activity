package api

import (
	"net/http"
	"student-information-system/internal/handlers"
)

// Router configure and return a new HTTP router.
func Router() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/api/endpoint", handlers.Endpoint())

	return mux
}
