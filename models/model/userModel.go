package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}