package http

import (
	"net/http"
)

func InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homepage)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/employees/", getEmployees)
	mux.HandleFunc("/departments/", getDepartments)
	return mux
}
