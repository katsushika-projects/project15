package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Thread struct {
	ID       string `gorm:"type:text;primaryKey"`
	Discript string `gorm:"column:discript"`
	ClassID  string `gorm:"not null;column:class_id"`
	Class    Class  `gorm:"foreignKey:ClassID;references:ID"`
}

func (u *Thread) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String() // UUIDを文字列として生成
	}
	return nil
}
