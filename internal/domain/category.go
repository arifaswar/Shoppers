package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100);not null unique" json:"name"`
	Products []Product `gorm:"foreignKey:CategoryID" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (category *Category) BeforeCreate(tx *gorm.DB) error {
	category.ID = uuid.New()
	return nil
}