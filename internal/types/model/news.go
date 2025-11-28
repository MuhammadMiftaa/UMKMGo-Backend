package model

import "time"

type News struct {
	ID          int        `json:"id" gorm:"primary_key"`
	Title       string     `json:"title" gorm:"type:varchar(255);not null"`
	Slug        string     `json:"slug" gorm:"type:varchar(255);not null;unique"`
	Excerpt     string     `json:"excerpt" gorm:"type:text"`
	Content     string     `json:"content" gorm:"type:text;not null"`
	Thumbnail   string     `json:"thumbnail" gorm:"type:text"`
	Category    string     `json:"category" gorm:"type:news_category;not null;default:'general'"`
	AuthorID    int        `json:"author_id" gorm:"not null"`
	IsPublished bool       `json:"is_published" gorm:"not null;default:false"`
	PublishedAt *time.Time `json:"published_at"`
	ViewsCount  int        `json:"views_count" gorm:"not null;default:0"`
	Base

	Author User      `json:"author" gorm:"foreignKey:AuthorID"`
	Tags   []NewsTag `json:"tags" gorm:"foreignKey:NewsID"`
}

type NewsTag struct {
	ID        int       `json:"id" gorm:"primary_key"`
	NewsID    int       `json:"news_id" gorm:"not null"`
	TagName   string    `json:"tag_name" gorm:"type:varchar(50);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"default:NOW()"`
}