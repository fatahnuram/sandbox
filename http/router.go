package http

import (
	"net/http"
)

func InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homepage)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/employees", listAllEmployees)
	mux.HandleFunc("/departments", listAllDepartments)
	return mux
}
