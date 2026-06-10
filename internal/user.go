package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID 	  uuid.UUID `gorm:"type:uuid;primary_key"`
	Name  string    `gorm:"type:varchar(100);not null"`
	Email string    `gorm:"type:varchar(100);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Role    string  `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}