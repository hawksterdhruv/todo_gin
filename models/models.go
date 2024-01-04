package models

import "time"

type Todo struct {
	Id          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tag struct {
}
type Comment struct {
}

type User struct {
}
