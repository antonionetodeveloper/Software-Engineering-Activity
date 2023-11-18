package handlers

import (
	"net/http"
)

// Endpoint ...
func Endpoint() http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte("Hello, World!"))
		},
	)
}
