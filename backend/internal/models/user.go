package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           string `gorm:"type:text;primaryKey"`
	Username     string `gorm:"not null;unique;column:username"`
	Password     string `gorm:"not null;column:password"`
	RefreshToken string `gorm:"column:refresh_token"`
}

type BlackList struct {
	AccessToken string `gorm:"column:access_token"`
}

// BeforeCreateフックでUUIDを生成
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String() // UUIDを文字列として生成
	}
	return nil
}
