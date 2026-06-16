package domain

import (
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type User struct {
	ID		uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Name	string         `gorm:"type:varchar(100);not null" json:"name"`
	Email	string         `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password string        `gorm:"type:varchar(255);not null" json:"-"`
	Addresses []Address      `gorm:"foreignKey:UserID" json:"addresses"`
	Role	string         `gorm:"type:varchar(50);not null:default:'user'" json:"role"`
	CreatedAt time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New()
	return nil
}