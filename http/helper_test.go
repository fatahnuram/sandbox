package http

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParsePathParameter(t *testing.T) {
	suites := []struct {
		name string
		path string
		want []string
	}{
		{name: "root, trailing", path: "/", want: []string{}},
		{name: "root, no trailing", path: "", want: []string{}},
		{name: "department, trailing", path: "/departments/", want: []string{"departments"}},
		{name: "department, no trailing", path: "/departments", want: []string{"departments"}},
		{name: "department id, trailing", path: "/departments/123/", want: []string{"departments", "123"}},
		{name: "department id, no trailing", path: "/departments/123", want: []string{"departments", "123"}},
		{name: "random, trailing", path: "/abce/def/gh/jklm/", want: []string{"abce", "def", "gh", "jklm"}},
		{name: "random, no trailing", path: "/abce/def/gh/jklm", want: []string{"abce", "def", "gh", "jklm"}},
	}

	for _, suite := range suites {
		t.Run(suite.name, func(t *testing.T) {
			got := parsePathParameter(suite.path)
			if len(got) != len(suite.want) {
				t.Errorf("result slice length different, want: %v, got: %v", len(suite.want), len(got))
			}
			if len(got) > 0 {
				for i := range got {
					if got[i] != suite.want[i] {
						t.Errorf("parsed path malformed, want: %v, got: %v", suite.want, got)
					}
				}
			}
		})
	}
}

func TestUnhandledRoute(t *testing.T) {
	suites := []struct {
		Name       string
		Url        string
		WantStatus int
		WantBody   ErrorMsg
	}{
		{
			Name:       "resource",
			Url:        "/employees/",
			WantStatus: http.StatusBadRequest,
			WantBody:   ErrorMsg{Error: false, Msg: UNHANDLED_ROUTE},
		},
		{
			Name:       "resource id",
			Url:        "/employees/123/",
			WantStatus: http.StatusBadRequest,
			WantBody:   ErrorMsg{Error: false, Msg: UNHANDLED_ROUTE},
		},
	}

	for _, suite := range suites {
		t.Run(suite.Name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, suite.Url, nil)
			if err != nil {
				t.Fatalf("failed to init http request: %v", err)
			}

			handleUnsupportedRoute(rec, req)

			resp := rec.Result()
			if resp.StatusCode != suite.WantStatus {
				t.Errorf("status code mismatch, want: %v, got: %v", suite.WantStatus, resp.StatusCode)
			}

			buffbody, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("failed to read http body: %v", err)
			}
			defer resp.Body.Close()

			msg := ErrorMsg{}
			err = json.Unmarshal(buffbody, &msg)
			if err != nil {
				t.Fatalf("failed to unmarshal json body: %v", err)
			}

			if msg.Error != suite.WantBody.Error {
				t.Errorf("resp body error flag mismatch, want: %v, got: %v", suite.WantBody.Error, msg.Error)
			}
			if msg.Msg != suite.WantBody.Msg {
				t.Errorf("resp body msg mismatch, want: %v, got: %v", suite.WantBody.Msg, msg.Msg)
			}
		})
	}
}
