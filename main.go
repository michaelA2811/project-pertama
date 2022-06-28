package main

import (
	"fmt"
	"project-pertama/entity"
	"project-pertama/service"
)

func main() {
	userService := service.NewUserService()

	if user, err := userService.Register(&entity.User{
		Username: "Mekel",
		Email:    "Mekel@kitabisa.com",
		Password: "321321",
		Age:      30,
	}); err != nil {
		fmt.Printf("Error when register user: %+v", err)
		return
	} else {
		fmt.Printf("Success register user: %+v", user)
	}

}
