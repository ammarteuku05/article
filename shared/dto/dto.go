package dto

import "time"

type RequestAuthor struct {
	Name string `json:"name" validate:"required"`
}

type RequestArticle struct {
	AuthorID string `json:"author_id" validate:"required,uuid4"`
	Title    string `json:"title" validate:"required,min=3,max=100"`
	Body     string `json:"body" validate:"required,min=10"`
}

type QueryGetArticle struct {
	AuthorName string `query:"author_name"`
	Titles     string `query:"title"`
	Body       string `query:"body"`
	Page       int    `query:"page"`
	PerPage    int    `query:"per_page"`
	Sort       string `query:"sort"`
}

type ResponseGetArticles struct {
	ArticleID  string    `json:"article_id"`
	AuthorID   string    `json:"author_id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	AuthorName string    `json:"author_name"`
	CreatedAt  time.Time `json:"created_at"`
}
