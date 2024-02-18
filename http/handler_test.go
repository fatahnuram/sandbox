package http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomepage(t *testing.T) {
	wantBody := "Welcome.\n"
	wantStatus := http.StatusOK

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("failed to initialize http request: %v\n", err)
	}

	homepage(rec, req)

	if rec.Result().StatusCode != wantStatus {
		t.Errorf("resp status not OK, want: %v, got: %v", wantStatus, rec.Result().StatusCode)
	}
	buffbody, err := io.ReadAll(rec.Result().Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v\n", err)
	}
	defer rec.Result().Body.Close()

	respBodyString := string(buffbody)
	if respBodyString != wantBody {
		t.Errorf("incorrect resp body, want: %v, got: %v", wantBody, respBodyString)
	}
}

func TestHealthz(t *testing.T) {
	wantBody := "ok\n"
	wantStatus := http.StatusOK

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/healthz", nil)
	if err != nil {
		t.Fatalf("failed to initialize http request: %v\n", err)
	}

	healthz(rec, req)

	resp := rec.Result()
	if resp.StatusCode != wantStatus {
		t.Errorf("resp status not OK, want: %v, got: %v", wantStatus, resp.StatusCode)
	}

	buffbody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v\n", err)
	}
	defer resp.Body.Close()

	body := string(buffbody)
	if body != wantBody {
		t.Errorf("incorrect resp body, want: %v, got: %v", wantBody, body)
	}
}
