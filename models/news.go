package models

type News struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
