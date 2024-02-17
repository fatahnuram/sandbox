package http

import (
	"encoding/json"
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

func listAllEmployees(resp http.ResponseWriter, req *http.Request) {
	employees, err := model.ListAllEmployees()
	if err != nil {
		log.Printf("[ERROR] %v", err)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(SOMETHING_WENT_WRONG))
	} else {
		emplbytes, err := json.Marshal(employees)
		if err != nil {
			log.Printf("[ERROR] %v", err)
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte(SOMETHING_WENT_WRONG))
		} else {
			resp.Header().Set("content-type", "application/json")
			resp.Write([]byte(emplbytes))
		}
	}
}

func listAllDepartments(resp http.ResponseWriter, req *http.Request) {
	depts, err := model.ListAllDepartments()
	if err != nil {
		log.Printf("[ERROR] %v", err)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(SOMETHING_WENT_WRONG))
	} else {
		deptbytes, err := json.Marshal(depts)
		if err != nil {
			log.Printf("[ERROR] %v", err)
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte(SOMETHING_WENT_WRONG))
		} else {
			resp.Write([]byte(deptbytes))
		}
	}
}
