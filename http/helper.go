package http

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type MsgPlaceholder struct {
	Msg string
}

type ErrorMsg struct {
	Error bool
	Msg   string
}

func parsePathParameter(path string) []string {
	// for root path, return empty slice with len == 0
	// handles both with trailing slash or not
	if path == "/" || len(path) == 0 {
		return []string{}
	}

	// remove slash at the beginning of path
	// e.g. /departments/123
	chars := strings.Split(path, "")
	if chars[0] == "/" {
		chars = chars[1:]
	}

	// remove trailing slash
	last := chars[len(chars)-1:]
	if last[0] == "/" {
		chars = chars[0 : len(chars)-1]
	}

	// construct a new path without trailing slash
	corepath := strings.Join(chars, "")

	// split based on slash character
	clean := strings.TrimSpace(corepath)
	s := strings.Split(clean, "/")

	return s
}

func sendResponse(writer http.ResponseWriter, statuscode int, payload interface{}) {
	writer.Header().Set("content-type", "application/json")

	_, ok := payload.(error)
	if ok {
		e := ErrorMsg{Error: true, Msg: SOMETHING_WENT_WRONG}
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(e)
		return
	}

	writer.WriteHeader(statuscode)
	err := json.NewEncoder(writer).Encode(payload)
	if err != nil {
		log.Printf("[ERROR] json marshal: %v", err)
		http.Error(writer, SOMETHING_WENT_WRONG, http.StatusInternalServerError)
	}
}

func wrapJsonResponseFromDB(writer http.ResponseWriter, err error, payload interface{}) {
	if err != nil {
		if err == sql.ErrNoRows {
			// no resources found
			log.Print("[WARN] not found")
			msg := ErrorMsg{Error: false, Msg: NOT_FOUND}
			sendResponse(writer, http.StatusNotFound, msg)
			return
		}

		// any unhandled errors
		log.Printf("[ERROR] server resp: %v", err)
		msg := ErrorMsg{Error: true, Msg: SOMETHING_WENT_WRONG}
		sendResponse(writer, http.StatusInternalServerError, msg)
		return
	}

	sendResponse(writer, http.StatusOK, payload)
}

func handleUnsupportedRoute(w http.ResponseWriter, r *http.Request) {
	log.Printf("[WARN] undhandled route: %v %v", r.Method, r.URL.Path)
	payload := ErrorMsg{Error: false, Msg: UNHANDLED_ROUTE}
	sendResponse(w, http.StatusBadRequest, payload)
}
