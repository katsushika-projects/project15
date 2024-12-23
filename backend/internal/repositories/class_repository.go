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

func (r *ClassRepository) FindById(id string) (*models.Class, error) {
	var class models.Class
	if err := r.db.Where("id = ?", id).First(&class).Error; err != nil {
		return nil, err
	}
	return &class, nil
}

func (r *ClassRepository) DeleteClass(class *models.Class) error {
	return r.db.Delete(class).Error
}

func (r *ClassRepository) FindClasses(group_id string) ([]*models.Class, error) {
	//該当するグループインスタンスを全て取得
	var classes []*models.Class
	if err := r.db.Where("group_id = ?", group_id).Find(&classes).Error; err != nil {
		return nil, err
	}

	if len(classes) == 0 {
		return nil, errors.New("groups don't exist")
	}

	return classes, nil
}
