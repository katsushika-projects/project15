package repositories

import (
	"errors"

	"github.com/moto340/project15/backend/internal/models"
	"gorm.io/gorm"
)

type ClassRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) *ClassRepository {
	return &ClassRepository{db: db}
}

func (r *ClassRepository) CreateClass(class *models.Class) error {
	return r.db.Create(class).Error
}

func (r *ClassRepository) FindByClass(classname, group_id string) error {
	var class models.Class
	if err := r.db.Where("classname = ? AND group_id = ?", classname, group_id).First(&class).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	return errors.New("class already exist")
}
