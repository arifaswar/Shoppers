package domain

import (
	"github.com/google/uuid"
	"time"
	"gorm.io/gorm"
)

type Product struct {
	ID		  uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
	Name	  string      `gorm:"type:varchar(100);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price	  float64     `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock	  int         `gorm:"not null" json:"stock"`
	CategoryID  uuid.UUID   `gorm:"type:uuid;not null" json:"category_id"`
	Category    Category    `gorm:"foreignKey:CategoryID" json:"category"`
	ImageURL  string      `gorm:"type:varchar(255)" json:"image_url"`
	CreatedAt  time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) error {
	product.ID = uuid.New()
	return nil
}