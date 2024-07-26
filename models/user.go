package main

import "time"

type User struct {
	ID             int       `gorm:"primaryKey" json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Password       string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Description    string    `json:"description"`
	Role           string    `json:"role"`
	ProfilePicture string    `json:"profile_picture"`
}
