package http

import (
	"log"
	"net/http"
	"strconv"

	model "github.com/fatahnuram/sandbox/db"
)

func homepage(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Welcome.\n"))
}

func healthz(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("ok\n"))
}

// listAllEmployees list or get employees based on request path
func getEmployees(resp http.ResponseWriter, req *http.Request) {
	path := parsePathParameter(req.URL.Path)
	if len(path) == 1 {
		log.Println("list employees")
		employees, err := model.ListAllEmployees()
		wrapJsonResponse(resp, err, employees)
	} else {
		log.Println("get employee by ID")

		param := path[1]
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Printf("[ERROR] convert to int: %v", err)
			msg := ErrorMsg{Error: true, Msg: INVALID_ID}
			sendResponse(resp, http.StatusBadRequest, msg)
			return
		}
		id64 := int64(id)

		empl, err := model.GetEmployeeById(id64)
		wrapJsonResponse(resp, err, empl)
	}
}

// listAllDepartments list or get departments based on request path
func getDepartments(resp http.ResponseWriter, req *http.Request) {
	path := parsePathParameter(req.URL.Path)
	if len(path) == 1 {
		log.Println("list departments")
		depts, err := model.ListAllDepartments()
		wrapJsonResponse(resp, err, depts)
	} else {
		log.Println("get department by ID")

		param := path[1]
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Printf("[ERROR] convert to int: %v", err)
			msg := ErrorMsg{Error: true, Msg: INVALID_ID}
			sendResponse(resp, http.StatusBadRequest, msg)
			return
		}
		id64 := int64(id)

		dep, err := model.GetDepartmentById(id64)
		wrapJsonResponse(resp, err, dep)
	}
}
