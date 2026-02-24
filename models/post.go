package models

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"not null;size:255" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	Excerpt   string    `gorm:"size:500" json:"excerpt"`
	AuthorID  uint      `gorm:"not null" json:"author_id"`
	Author    User      `gorm:"foreignKey:AuthorID" json:"author"`
	Category  string    `gorm:"size:100" json:"category"`
	Tags      string    `gorm:"size:255" json:"tags"`
	Published bool      `gorm:"default:true" json:"published"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
