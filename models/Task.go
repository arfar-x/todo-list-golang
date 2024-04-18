package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model `json:"-"` // To hide gorm specific fields.
	ID         uint       `gorm:"primaryKey" json:"id" flag:"id"`
	Name       string     `json:"name" flag:"name"`
	Done       bool       `json:"done" flag:"done"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
