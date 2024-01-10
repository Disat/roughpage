package model

import (
	"html/template"
	"time"
)

type ArticleConfig struct {
	Title    string
	Tags     []string
	Category string
	Preview  string
	Top      bool
	Toc      bool
	Markdown string
}
type Article struct {
	ArticleConfig
	CreateTime       time.Time
	CreateTimeString string
	UpdateTime       time.Time
	UpdateTimeString string
	ContentHTML      template.HTML
	Filename         string
	CreateYear       int
	CreateMonth      int
}
