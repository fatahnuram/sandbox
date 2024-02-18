package http

import (
	"log"
	"net/http"

	model "github.com/fatahnuram/sandbox/db"
)

func homepage(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Welcome.\n"))
}

func healthz(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("ok\n"))
}

// listAllEmployees list or get employees based on request path
func listAllEmployees(resp http.ResponseWriter, req *http.Request) {
	path := parsePathParameter(req.URL.Path)
	if len(path) == 1 {
		log.Println("list employees")
		employees, err := model.ListAllEmployees()
		wrapJsonResponse(resp, err, employees)
	} else {
		log.Println("get employee by ID")
		wrapJsonResponse(resp, nil, MsgPlaceholder{Msg: "wip"})
	}
}

// listAllDepartments list or get departments based on request path
func listAllDepartments(resp http.ResponseWriter, req *http.Request) {
	path := parsePathParameter(req.URL.Path)
	if len(path) == 1 {
		log.Println("list departments")
		depts, err := model.ListAllDepartments()
		wrapJsonResponse(resp, err, depts)
	} else {
		log.Println("get department by ID")
		wrapJsonResponse(resp, nil, MsgPlaceholder{Msg: "wip"})
	}
}
