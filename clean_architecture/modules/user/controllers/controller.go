package controllers

import (
	"eiei/clean_architecture/entities"
)

type userController struct {
	userUsecase entities.UserUsecase
}

func NewUserController(userUsecase entities.UserUsecase) *userController {
	return &userController{
		userUsecase: userUsecase,
	}
}
