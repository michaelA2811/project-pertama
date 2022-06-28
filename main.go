package main

import (
	"project-pertama/entity"
	"project-pertama/service"
	"time"
)

func main() {
	userService := service.NewUserService()

	userService.Register(&entity.User{
		Id:        1,
		Username:  "Mekel",
		Email:     "Mekel@kitabisa.com",
		Password:  "321321",
		Age:       30,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

}
