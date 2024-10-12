package model

type Response struct {
	Result string `json:"result"`
	Title  string `json:"title"`
	Day    string `json:"day"`
	Tag    string `json:"tag"`
}
