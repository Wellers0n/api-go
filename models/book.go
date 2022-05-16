package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float32        `json:"price"`
	Author      string         `json:"author"`
	ImageUrl    string         `json:"img_url"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
}
