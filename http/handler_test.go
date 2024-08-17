package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomepage(t *testing.T) {
	wantBody := MsgPlaceholder{Msg: "Welcome."}
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

	var msg MsgPlaceholder
	err = json.NewDecoder(rec.Result().Body).Decode(&msg)
	if err != nil {
		t.Fatalf("failed to decode body: %v\n", err)
	}

	if msg.Msg != wantBody.Msg {
		t.Errorf("incorrect resp body, want: %v, got: %v", wantBody.Msg, msg.Msg)
	}
}

func TestHealthz(t *testing.T) {
	wantBody := MsgPlaceholder{Msg: "ok"}
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

	var msg MsgPlaceholder
	err = json.NewDecoder(resp.Body).Decode(&msg)
	if err != nil {
		t.Fatalf("failed to decode body: %v\n", err)
	}

	if msg.Msg != wantBody.Msg {
		t.Errorf("incorrect resp body, want: %v, got: %v", wantBody.Msg, msg.Msg)
	}
}
