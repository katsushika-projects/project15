package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"type:text;primaryKey"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

// BeforeCreateフックでUUIDを生成
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String() // UUIDを文字列として生成
	}
	return nil
}
