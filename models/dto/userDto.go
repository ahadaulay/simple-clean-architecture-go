package dto

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64	`json:"id"`
	Name      string	`json:"name"`
	Email     string	`json:"email"`
	Password  string	`json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type UserCreate struct{
	Name      string	`json:"name"`
	Email     string	`json:"email"`
	Password  string	`json:"password"`
}