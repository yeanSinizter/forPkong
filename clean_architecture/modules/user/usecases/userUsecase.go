package usecases

import (
	"eiei/clean_architecture/entities"
)

func (u *userUsecase) AddUser(user *entities.UserModel) (*entities.UserModel, error) {
	return u.userRepo.InsertUser(user)
}
