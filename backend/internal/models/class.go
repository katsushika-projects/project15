package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	ID        string `gorm:"type:text;primaryKey"`
	ClassName string `gorm:"not null;column:classname"`
	GroupID   string `gorm:"not null;column:group_id"`
	Group     Group  `gorm:"foreignKey:GroupID;references:ID"`
}

func (u *Class) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String() // UUIDを文字列として生成
	}
	return nil
}
