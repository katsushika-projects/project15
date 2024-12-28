package repositories

import (
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
