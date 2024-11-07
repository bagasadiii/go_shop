package model

import "time"

type User struct {
	UserID		string		`gorm:"primaryKey"`
	Username	string		`gorm:"unique, not null"`
	Email		string		`gorm:"unique, not null"`
	Password	string		`gorm:"not null"`
	CreatedAt	time.Time
}