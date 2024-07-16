package http

import (
	"log"
	"net/http"
	"strconv"

	model "github.com/fatahnuram/sandbox/db"
)

func notFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 not found.\n"))
	})
}

func homepage(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Welcome.\n"))
}

func healthz(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("ok\n"))
}

// handleEmployees handle employees CRUD based on request method and path
func handleEmployees(resp http.ResponseWriter, req *http.Request) {
	path := parsePathParameter(req.URL.Path)
	if len(path) == 1 {
		switch req.Method {
		case http.MethodGet:
			log.Println("list employees")
			employees, err := model.ListAllEmployees()
			wrapJsonResponse(resp, err, employees)

		default:
			handleUnsupportedRoute(resp, req)
		}
	} else {
		param := path[1]
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Printf("[ERROR] convert to int: %v", err)
			msg := ErrorMsg{Error: true, Msg: INVALID_ID}
			sendResponse(resp, http.StatusBadRequest, msg)
			return
		}
		id64 := int64(id)

		switch req.Method {
		case http.MethodGet:
			log.Println("get employee by ID")
			empl, err := model.GetEmployeeById(id64)
			wrapJsonResponse(resp, err, empl)

		case http.MethodDelete:
			log.Println("delete employee by ID")
			rows, err := model.DeleteEmployeeById(id64)
			if err != nil {
				log.Printf("[ERROR] server resp: %v", err)
				msg := ErrorMsg{Error: true, Msg: SOMETHING_WENT_WRONG}
				sendResponse(resp, http.StatusInternalServerError, msg)
				return
			}
			if rows == 0 {
				log.Print("[WARN] not found")
				msg := ErrorMsg{Error: false, Msg: NOT_FOUND}
				sendResponse(resp, http.StatusNotFound, msg)
				return
			}
			msg := MsgPlaceholder{Msg: RESOURCE_DELETED}
			sendResponse(resp, http.StatusOK, msg)

		default:
			handleUnsupportedRoute(resp, req)
		}
	}
}

// handleDepartments handle departments CRUD based on request method and path
func handleDepartments(resp http.ResponseWriter, req *http.Request) {
	path := parsePathParameter(req.URL.Path)
	if len(path) == 1 {
		switch req.Method {
		case http.MethodGet:
			log.Println("list departments")
			depts, err := model.ListAllDepartments()
			wrapJsonResponse(resp, err, depts)

		default:
			handleUnsupportedRoute(resp, req)
		}
	} else {
		param := path[1]
		id, err := strconv.Atoi(param)
		if err != nil {
			log.Printf("[ERROR] convert to int: %v", err)
			msg := ErrorMsg{Error: true, Msg: INVALID_ID}
			sendResponse(resp, http.StatusBadRequest, msg)
			return
		}
		id64 := int64(id)

		switch req.Method {
		case http.MethodGet:
			log.Println("get department by ID")
			dep, err := model.GetDepartmentById(id64)
			wrapJsonResponse(resp, err, dep)

		case http.MethodDelete:
			log.Println("delete department by ID")
			rows, err := model.DeleteDepartmentById(id64)
			if err != nil {
				log.Printf("[ERROR] server resp: %v", err)
				msg := ErrorMsg{Error: true, Msg: SOMETHING_WENT_WRONG}
				sendResponse(resp, http.StatusInternalServerError, msg)
				return
			}
			if rows == 0 {
				log.Print("[WARN] not found")
				msg := ErrorMsg{Error: false, Msg: NOT_FOUND}
				sendResponse(resp, http.StatusNotFound, msg)
				return
			}
			msg := MsgPlaceholder{Msg: RESOURCE_DELETED}
			sendResponse(resp, http.StatusOK, msg)

		default:
			handleUnsupportedRoute(resp, req)
		}

	}
}
