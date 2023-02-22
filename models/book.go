package models

type Book struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"type:varchar(300)" json:"title"`
	Description string `gorm:"text" json:"description"`
	Author      string `gorm:"varchar(300)" json:"author"`
	PublishDate string `gorm:"data" json:"publish_date"`
}
