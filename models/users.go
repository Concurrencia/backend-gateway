package models

import (
	"apigo/db"
)

// User
type User struct {
	ID            string         `json:"id"`
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	Consultations []Consultation `json:"consultations"`
}

// RequestUserDto
type RequestUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsersResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
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
