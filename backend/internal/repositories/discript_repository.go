package repositories

import (
	"errors"

	"github.com/moto340/project15/backend/internal/models"
	"gorm.io/gorm"
)

type DiscriptRepository struct {
	db *gorm.DB
}

func NewDiscriptRepository(db *gorm.DB) *DiscriptRepository {
	return &DiscriptRepository{db: db}
}

func (r *DiscriptRepository) CreateDiscript(discript *models.Thread) error {
	return r.db.Create(discript).Error
}

func (r *DiscriptRepository) GetDiscripts(class_id string) ([]*models.Thread, error) {
	var discripts []*models.Thread
	if err := r.db.Where("class_id = ?", class_id).Find(&discripts).Error; err != nil {
		return nil, err
	}

	if len(discripts) == 0 {
		return nil, errors.New("discripts don't exist")
	}

	return discripts, nil
}
