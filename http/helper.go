package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type MsgPlaceholder struct {
	Msg string
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
	chars = chars[1:]

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

func wrapJsonResponse(writer http.ResponseWriter, err error, payload ...interface{}) {
	if err != nil {
		log.Printf("[ERROR] %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(SOMETHING_WENT_WRONG))
	} else {
		payloadbytes, err := json.Marshal(payload)
		if err != nil {
			log.Printf("[ERROR] %v", err)
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(SOMETHING_WENT_WRONG))
		} else {
			writer.Header().Set("content-type", "application/json")
			writer.Write([]byte(payloadbytes))
		}
	}
}
