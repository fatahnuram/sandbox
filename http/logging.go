package http

import (
	"fmt"
	"net/http"
	"time"
)

func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		start := time.Now()
		next.ServeHTTP(w, r.WithContext(ctx))
		status, ok := r.Context().Value(CTX_HTTP_RESP_STATUS).(int)
		if ok {
			fmt.Printf("%s %d %s %v\n", r.Method, status, r.URL.Path, time.Since(start))
		} else {
			fmt.Printf("%s %s %v\n", r.Method, r.URL.Path, time.Since(start))
		}
	})
}
