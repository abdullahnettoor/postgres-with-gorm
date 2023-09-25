package main

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary key"`
	FirstName   string    `gorm:"not null"`
	LastName    string    `gorm:"default:null"`
	Email       string    `gorm:"unique;not null"`
	Password    string    `gorm:"not null"`
	DateOfBirth time.Time `gorm:"type:date;not null"`
}
