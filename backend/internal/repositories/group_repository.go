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

func (r *GroupRepository) DeleteGroup(group *models.Group) error {
	return r.db.Delete(group).Error
}

func (r *GroupRepository) FindByGroup(university, fculty, department, grade string) error {
	//グループ作成時にグループが既存か検証
	var group models.Group
	if err := r.db.Where("university = ? AND fculty = ? AND department = ? AND grade = ?", university, fculty, department, grade).First(&group).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			return err
		}
	}

	return errors.New("group already exist")
}

func (r *GroupRepository) FindById(id string) (*models.Group, error) {
	//idに該当する特定のグループインスタンスを取得
	var group models.Group
	if err := r.db.Where("id = ?", id).First(&group).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *GroupRepository) FindGroup(university, fculty, department, grade string) ([]*models.Group, error) {
	//該当するグループインスタンスを全て取得
	var groups []*models.Group
	if err := r.db.Where("university = ? AND fculty = ? AND department = ? AND grade = ?", university, fculty, department, grade).Find(&groups).Error; err != nil {
		return nil, err
	}

	if len(groups) == 0 {
		return nil, errors.New("groups don't exist")
	}

	return groups, nil
}
