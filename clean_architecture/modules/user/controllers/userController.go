package controllers

import (
	"eiei/clean_architecture/entities"

	"github.com/labstack/echo/v4"
)

func (u *userController) HHHHHHH(ctx echo.Context) error {
	user := new(entities.UserModel)
	err := ctx.Bind(&user)

	if err != nil {
		return ctx.JSON(400, map[string]interface{}{
			"error": err.Error(),
		})
	}

	userRes, err := u.userUsecase.AddUser(user)

	if err != nil {
		return ctx.JSON(500, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, userRes)

}
