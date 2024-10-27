package model

type MorphRequest struct {
	AppID      string `json:"app_id"`
	RequestID  string `json:"request_id,omitempty"`
	Sentence   string `json:"sentence"`
	InfoFilter string `json:"info_filter,omitempty"`
}

type MorphResponse struct {
	InfoFilter string        `json:"info_filter"`
	RequestID  string        `json:"request_id"`
	WordList   [][][]string  `json:"word_list"`
}
