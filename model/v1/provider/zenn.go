package model

type Article struct {
	ID               int    `json:"id"`
	PostType         string `json:"post_type"`
	Title            string `json:"title"`
	Slug             string `json:"slug"`
	CommentsCount    int    `json:"comments_count"`
	LikedCount       int    `json:"liked_count"`
	BodyLettersCount int    `json:"body_letters_count"`
	ArticleType      string `json:"article_type"`
	Emoji            string `json:"emoji"`
	PublishedAt      string `json:"published_at"`
	Path             string `json:"path"`
	User             User   `json:"user"`
}

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	AvatarSmallURL string `json:"avatar_small_url"`
}
