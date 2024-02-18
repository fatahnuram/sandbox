package http

import (
	"fmt"
	"io"
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
		WantBody   string
	}{
		{Name: "get root", Url: fmt.Sprintf("%v/", srv.URL), WantStatus: http.StatusOK, WantBody: "Welcome.\n"},
		{Name: "get healthz", Url: fmt.Sprintf("%v/healthz", srv.URL), WantStatus: http.StatusOK, WantBody: "ok\n"},
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

			buff, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("failed to read body: %v", err)
			}
			defer resp.Body.Close()

			body := string(buff)
			if body != suite.WantBody {
				t.Errorf("incorrect resp body, want: %v, got: %v", suite.WantBody, body)
			}
		})
	}
}
