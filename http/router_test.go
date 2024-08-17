package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterGet(t *testing.T) {
	srv := httptest.NewServer(InitRoutes())
	defer srv.Close()

	suites := []struct {
		Name       string
		Url        string
		WantStatus int
		WantBody   MsgPlaceholder
	}{
		{
			Name:       "get root",
			Url:        fmt.Sprintf("%v/", srv.URL),
			WantStatus: http.StatusOK,
			WantBody:   MsgPlaceholder{Msg: "Welcome."},
		},
		{
			Name:       "get healthz",
			Url:        fmt.Sprintf("%v/healthz", srv.URL),
			WantStatus: http.StatusOK,
			WantBody:   MsgPlaceholder{Msg: "ok"},
		},
	}

	for _, suite := range suites {
		t.Run(suite.Name, func(t *testing.T) {
			resp, err := http.Get(suite.Url)
			if err != nil {
				t.Fatalf("failed to request http: %v", err)
			}

			if resp.StatusCode != suite.WantStatus {
				t.Errorf("resp status not match, want: %v, got: %v", suite.WantStatus, resp.StatusCode)
			}

			var body MsgPlaceholder
			err = json.NewDecoder(resp.Body).Decode(&body)
			if err != nil {
				t.Fatalf("failed to decode body: %v", err)
			}

			if body.Msg != suite.WantBody.Msg {
				t.Errorf("incorrect resp body, want: %v, got: %v", suite.WantBody.Msg, body.Msg)
			}
		})
	}
}
