package service

import (
	"errors"
	"project-pertama/entity"
)

type UserServiceIface interface {
	Register(user *entity.User) (*entity.User, error)
}

type userService struct{}

func NewUserService() UserServiceIface {
	return &userService{}
}

func (u *userService) Register(user *entity.User) (*entity.User, error) {
	x := 1
	if x == 1 {
		return nil, errors.New("eror bang")
	}
	return user, nil
}
