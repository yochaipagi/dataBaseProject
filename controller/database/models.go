package database

import (
	"time"
)

// common fields to the database tables
type Base struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Article struct {
	Base
	PublishedAt  time.Time     `json:"published_at"`
	Author       string        `json:"author"`
	Title        string        `json:"title"`
	Source       string        `json:"source"`
	ArticlePages []ArticlePage `json:",omitempty"` // one to many relationship (article have many pages)
	PagesCount   int           `json:"pages_count"`
}

type ArticlePage struct {
	Base
	ArticleID    uint          `gorm:"index"`
	PageNumber   int           `gorm:"index"`
	ArticleLines []ArticleLine `json:",omitempty"` //one to many relationship (page have many lines)
}

type ArticleLine struct {
	Base
	ArticlePageID uint          `gorm:"index"`
	LineNumber    int           `gorm:"index"`
	ArticleWords  []ArticleWord `json:",omitempty"` //one to many relationship (line have many words)
	WordCount     int
}

type ArticleWord struct {
	Base
	ArticleLineID uint `gorm:"index"`
	WordNumber    int  `gorm:"index"`
	Word          string
	CharCount     int
}

type WordGroup struct {
	Base
	Name  string `json:"name" binding:"required"`
	Words []Word `json:"words" binding:"required"`
}

type Word struct {
	Base
	WordGroupID uint   `gorm:"index"`
	Word        string `json:"word" binding:"required"` //non empty must be represented in the json.
}

type LinguisticExpr struct {
	Base
	Expression string `json:"expression" binding:"required"`
}

/* rename tables from gorm default*/

func (a ArticleLine) TableName() string {
	return "article_lines"
}

func (a ArticleWord) TableName() string {
	return "article_words"
}

func (a ArticlePage) TableName() string {
	return "article_pages"
}

func (a Word) TableName() string {
	return "words"
}
