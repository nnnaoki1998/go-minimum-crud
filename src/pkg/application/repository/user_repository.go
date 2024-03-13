package repository

import (
	"go-minimum-crud/src/pkg/domain/model"
)

type UserRepository interface {
	SelectAll() ([]model.User, error)
	Select(id model.UserId) (*model.User, error)
	Insert(user model.NewUser) error
	Update(id model.UserId, user model.NewUser) error
	Delete(id model.UserId) error
}
