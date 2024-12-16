package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	ID         string `gorm:"type:text;primaryKey"`
	University string `gorm:"not null;column:university"`
	Fculty     string `gorm:"not null;column:fculty"`
	Department string `gorm:"not null;column:department"`
	Grade      string `gorm:"not null;column:grade"`
}

func (u *Group) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String() // UUIDを文字列として生成
	}
	return nil
}
