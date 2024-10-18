package notes_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	h "github.com/yashikota/chronotes/api/v1/debug"
	"github.com/yashikota/chronotes/pkg/notes"
)

func TestCronHandler(t *testing.T) {
	// Cron関数の実行
	notes.Cron()

	// healthエンドポイントを叩く
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// レコーダーを作成
	w := httptest.NewRecorder()

	// ハンドラーを呼び出す
	h.HealthHandler(w, req)

	// レスポンスの検証
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status OK; got %v", w.Code)
	}

	expected := `{"message":"pong"}`
	actual := strings.TrimSpace(w.Body.String())
	if actual != expected {
		t.Fatalf("Expected body %v; got %v", expected, actual)
	}
}
