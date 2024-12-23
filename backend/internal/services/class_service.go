package services

import (
	"github.com/moto340/project15/backend/internal/models"
	"github.com/moto340/project15/backend/internal/repositories"
)

type ClassService struct {
	classService *repositories.ClassRepository
}

func NewClassService(classRepo *repositories.ClassRepository) *ClassService {
	return &ClassService{classService: classRepo}
}

func (s *ClassService) CreateClass(classname, group_id string) error {
	//classの重複確認
	if err := s.classService.FindByClass(classname, group_id); err != nil {
		return err
	}

	class := models.Class{
		ClassName: classname,
		GroupID:   group_id,
	}

	if err := s.classService.CreateClass(&class); err != nil {
		return err
	}

	return nil
}

func (s *ClassService) DeleteClass(id string) error {
	class, err := s.classService.FindById(id)
	if err != nil {
		return err
	}

	if err1 := s.classService.DeleteClass(class); err1 != nil {
		return err1
	}

	return nil
}
