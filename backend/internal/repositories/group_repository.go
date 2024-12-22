package repositories

import (
	"errors"

	"github.com/moto340/project15/backend/internal/models"
	"gorm.io/gorm"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (r *GroupRepository) CreateGroup(group *models.Group) error {
	return r.db.Create(group).Error
}

func (r *GroupRepository) FindByGroup(university string, fculty string, department string, grade string) error {
	var group models.Group
	if err := r.db.Where("university = ? AND fculty = ? AND department = ? AND grade = ?", university, fculty, department, grade).First(&group).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			return err
		}
	}
	return errors.New("group already exits")
}
