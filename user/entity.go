package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int    `gorm:"size:36; not null; uniqueIndex; primary_key"`
	Name      string `gorm:"size:100; not null"`
	Email     string `gorm:"size:100; not null; uniqueIndex"`
	Password  string `gorm:"size:100; not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
