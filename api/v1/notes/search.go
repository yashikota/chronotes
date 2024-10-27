package notes

import (
	"log/slog"
	"net/http"

	"github.com/yashikota/chronotes/model/v1"
	n "github.com/yashikota/chronotes/pkg/notes"
	"github.com/yashikota/chronotes/pkg/users"
	"github.com/yashikota/chronotes/pkg/utils"
)

// SearchHandler - ユーザーのノートを検索して、指定された単語を含むノートを返す
func SearchNoteHandler(w http.ResponseWriter, r *http.Request) {
	// トークンの検証
	user := model.NewUser()
	user.UserID = r.Context().Value(utils.TokenKey).(utils.Token).UserID

	// トークンが存在するか確認
	key := "jwt:" + user.UserID
	if _, err := utils.GetToken(key); err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Token validation passed")

	// リクエストパラメータから検索ワードを取得
	word, err := utils.GetQueryParam(r, "word", true)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	slog.Info("Search word received", "word", word)

	// ユーザー情報の取得
	user, err = users.GetUser(user)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusBadRequest, err)
		return
	}

	// ノートを検索
	matchingNotes, err := n.SearchNote(user.UserID, word)
	if err != nil {
		utils.ErrorJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("Notes search completed")

	// レスポンスの作成と送信
	utils.SuccessJSONResponse(w, matchingNotes)
}
