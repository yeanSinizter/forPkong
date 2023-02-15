package usecases

import (
	"eiei/clean_architecture/entities"
)

type userUsecase struct {
	userRepo entities.UserRepositories
}

func NewUserUsecase(userRepo entities.UserRepositories) *userUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
