package services

import (
	"github.com/moto340/project15/backend/internal/models"
	"github.com/moto340/project15/backend/internal/repositories"
)

type ClassService struct {
	classRepository *repositories.ClassRepository
}

func NewClassService(classRepo *repositories.ClassRepository) *ClassService {
	return &ClassService{classRepository: classRepo}
}

func (s *ClassService) CreateClass(classname, group_id string) error {
	//classの重複確認
	if err := s.classRepository.FindByClass(classname, group_id); err != nil {
		return err
	}

	class := models.Class{
		ClassName: classname,
		GroupID:   group_id,
	}

	if err := s.classRepository.CreateClass(&class); err != nil {
		return err
	}

	return nil
}

func (s *ClassService) DeleteClass(id string) error {
	class, err := s.classRepository.FindById(id)
	if err != nil {
		return err
	}

	if err1 := s.classRepository.DeleteClass(class); err1 != nil {
		return err1
	}

	return nil
}

func (s *ClassService) GetClass(id string) (*models.Class, error) {
	class, err := s.classRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return class, nil
}

func (s *ClassService) GetClasses(group_id string) ([]*models.Class, error) {
	classes, err := s.classRepository.FindClasses(group_id)
	if err != nil {
		return nil, err
	}

	return classes, nil
}
