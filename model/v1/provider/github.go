package model

type CommitInfo struct {
	Message string       `json:"message"`
	Changes []FileChange `json:"changes"`
	Period  string       `json:"period"`
}

type FileChange struct {
	Filename  string `json:"filename"`
	Status    string `json:"status"`
	Additions int    `json:"additions"`
	Deletions int    `json:"deletions"`
	Changes   int    `json:"changes"`
	Patch     string `json:"patch"`
}
