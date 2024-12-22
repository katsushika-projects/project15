package services

import (
	"github.com/moto340/project15/backend/internal/models"
	"github.com/moto340/project15/backend/internal/repositories"
)

type GroupService struct {
	groupRepository *repositories.GroupRepository
}

func NewGroupService(groupRepo *repositories.GroupRepository) *GroupService {
	return &GroupService{groupRepository: groupRepo}
}

func (s *GroupService) CreateGroup(university, fculty, department, grade string) error {
	// グループの重複確認
	if err := s.groupRepository.FindByGroup(university, fculty, department, grade); err != nil {
		return err
	}

	group := models.Group{
		University: university,
		Fculty:     fculty,
		Department: department,
		Grade:      grade,
	}
	if err := s.groupRepository.CreateGroup(&group); err != nil {
		return err
	}

	return nil
}

func (s *GroupService) DeleteGroup(id string) error {
	group, err := s.groupRepository.FindById(id)
	if err != nil {
		return err
	}

	if err1 := s.groupRepository.DeleteGroup(group); err1 != nil {
		return err1
	}

	return nil
}
