package common

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID         uint           `json:"id"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
