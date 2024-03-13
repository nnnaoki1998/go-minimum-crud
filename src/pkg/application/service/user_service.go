package service

import (
	"go-minimum-crud/src/pkg/application/repository"
	"go-minimum-crud/src/pkg/domain/model"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (s *UserService) Index() ([]model.User, error) {
	return s.UserRepository.SelectAll()
}

func (s *UserService) Show(id model.UserId) (*model.User, error) {
	return s.UserRepository.Select(id)
}

func (s *UserService) Create(user model.NewUser) error {
	return s.UserRepository.Insert(user)
}

func (s *UserService) Update(id model.UserId, user model.NewUser) error {
	return s.UserRepository.Update(id, user)
}

func (s *UserService) Delete(id model.UserId) error {
	return s.UserRepository.Delete(id)
}
