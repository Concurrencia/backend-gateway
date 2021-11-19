package models

import (
	"apigo/db"
)

// User
type User struct {
	ID            int64          `json:"id"`
	Username      string         `json:"username"`
	Password      string         `json:"password"`
	Email         string         `json:"email"`
	Consultations []Consultation `json:"consultations"`
}

// RequestUserDto
type RequestUserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// UserLogin
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Users
type Users []User

func MigrarUser() {
	db.Database.AutoMigrate(User{})
}
