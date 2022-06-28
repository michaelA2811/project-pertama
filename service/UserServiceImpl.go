package service

import (
	"fmt"
	"project-pertama/entity"
)

type UserServiceIface interface {
	Register(user *entity.User)
}

type userService struct{}

func NewUserService() UserServiceIface {
	return &userService{}
}

func (u *userService) Register(user *entity.User) {
	fmt.Println(user)
}
