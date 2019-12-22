package models

import (
	"time"
)

// Article represent the article model
type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Author    Author    `json:"author"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type ArticleDTO struct {
	ID        int64
	Title     string
	Content   string
	Author    Author
	UpdatedAt string
	CreatedAt string
}

const (
	sqliteTimeFormat = "2017-05-18 13:50:19"
)

func (a *ArticleDTO) ToArticle() *Article {
	up, _ := time.Parse(sqliteTimeFormat, a.UpdatedAt)
	ca, _ := time.Parse(sqliteTimeFormat, a.CreatedAt)

	return &Article{
		ID:        a.ID,
		Title:     a.Title,
		Content:   a.Content,
		Author:    a.Author,
		UpdatedAt: up,
		CreatedAt: ca,
	}
}
