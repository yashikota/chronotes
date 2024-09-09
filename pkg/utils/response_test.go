package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccessJsonResponse(t *testing.T) {
	w := httptest.NewRecorder()

	// テストデータ
	testData := map[string]string{"key": "value"}

	SuccessJsonResponse(w, testData)

	// ステータスコードを確認
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Content-Typeヘッダーを確認
	contentType := w.Header().Get("Content-Type")
	expectedContentType := "application/json; charset=utf-8"
	if contentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, got %s", expectedContentType, contentType)
	}

	// レスポンスボディを確認
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	if response["key"] != "value" {
		t.Errorf("Expected response body to contain {'key': 'value'}, got %v", response)
	}
}

func TestErrorJsonResponse(t *testing.T) {
	w := httptest.NewRecorder()

	// テストデータ
	testError := errors.New("test error")
	testStatus := http.StatusBadRequest

	// 関数を呼び出し
	ErrorJsonResponse(w, testStatus, testError)

	// ステータスコードを確認
	if w.Code != testStatus {
		t.Errorf("Expected status code %d, got %d", testStatus, w.Code)
	}

	// Content-Typeヘッダーを確認
	contentType := w.Header().Get("Content-Type")
	expectedContentType := "application/json; charset=utf-8"
	if contentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, got %s", expectedContentType, contentType)
	}

	// レスポンスボディを確認
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	if response["message"] != testError.Error() {
		t.Errorf("Expected response body to contain {'message': '%s'}, got %v", testError.Error(), response)
	}
}
