package models

import "time"

type Post struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint `gorm:"not null"`
	User        User
	Caption     string `gorm:"unique;not null"`
	PicturePath string `gorm:"unique;not null"`
	createdAt   time.Time
}
