package services

import (
	"github.com/moto340/project15/backend/internal/models"
	"github.com/moto340/project15/backend/internal/repositories"
)

type DiscriptService struct {
	discriptRepository *repositories.DiscriptRepository
}

func NewDiscriptService(discriptRepo *repositories.DiscriptRepository) *DiscriptService {
	return &DiscriptService{discriptRepository: discriptRepo}
}

func (s *DiscriptService) CreateDiscript(discript, class_id string) error {
	discription := models.Thread{
		Discript: discript,
		ClassID:  class_id,
	}

	if err := s.discriptRepository.CreateDiscript(&discription); err != nil {
		return err
	}

	return nil
}

func (s *DiscriptService) GetDiscripts(class_id string) ([]*models.Thread, error) {
	discripts, err := s.discriptRepository.GetDiscripts(class_id)
	if err != nil {
		return nil, err
	}

	return discripts, nil
}
