package db

type User struct {
	ID       string `gorm:"primaryKey"`
	Name     string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type Note struct {
	ID      string `gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    string `json:"tags"`
}

type Account struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Provider string `json:"provider"`
}
