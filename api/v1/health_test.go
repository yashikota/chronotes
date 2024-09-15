package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	h "github.com/yashikota/chronotes/api/v1"
)

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(h.HealthHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v, expected %v", status, http.StatusOK)
	}

	expected := `{"message":"pong"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("Unexpected response body: got %v, expected %v", rr.Body.String(), expected)
	}
}
