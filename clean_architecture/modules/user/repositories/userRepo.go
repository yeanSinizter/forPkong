package repositories

import (
	"eiei/clean_architecture/entities"
)

func (r *userRepo) InsertUser(user *entities.UserModel) (*entities.UserModel, error) {
	result := r.db.Create(&user).Error

	return user, result
}
