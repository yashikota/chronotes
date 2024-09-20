package debug

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yashikota/chronotes/pkg/utils"
)

func FakeHandler(w http.ResponseWriter, r *http.Request) {
	data := "2024-09-21T09:00:00Z"
	content := `今日は、ログイン機能を実装した。

- メールアドレスとパスワードの入力を受け付け、データベースに登録されたユーザーと照合する。
- パスワードが一致すれば、JWTトークンを生成し、Redisに保存する。
- トークンは、ユーザーのIDと名前を含んでいる。
- Redisにトークンを保存することで、ユーザーの認証情報を保持し、次回のログイン時に再利用できるようにする。

この変更により、ユーザーはログインして、Chronotesの機能を利用できるようになる。

また、Redisの接続設定も変更した。
- Docker Composeの設定に Redis Commander を追加し、Redisの管理を容易にした。
- Nginxの設定も変更して、Redis Commanderにアクセスできるようにした。

ログイン機能の実装は順調に進んでいる。
今後、ログアウト機能やユーザーの情報を取得する機能などを追加していく予定だ。
`
    html := utils.Md2HTML([]byte(content))

    // Create the structure
    structure := struct {
        Data    string `json:"data"`
        Content string `json:"content"`
    }{
        Data:    data,
        Content: string(html),
    }

    // カスタムエンコーダーを作成
    buffer := &bytes.Buffer{}
    encoder := json.NewEncoder(buffer)
    encoder.SetEscapeHTML(false)
    encoder.SetIndent("", "  ")

    // 構造体をJSONエンコード
    if err := encoder.Encode(structure); err != nil {
        fmt.Println("Error encoding JSON:", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // エンコードされたJSONを取得
    jsonData := buffer.Bytes()

    // レスポンスの作成
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}
