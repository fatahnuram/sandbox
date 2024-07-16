package http

import (
	"net/http"
)

func InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", notFoundHandler())
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/employees/", handleEmployees)
	mux.HandleFunc("/departments/", handleDepartments)
	return withLogging(mux)
}
