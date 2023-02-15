package routes

import (
	"eiei/clean_architecture/modules/user/controllers"
	"eiei/clean_architecture/modules/user/repositories"
	"eiei/clean_architecture/modules/user/usecases"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(e *echo.Echo, dbPg *gorm.DB) {
	// db, _ := db.InitSql()
	userRepo := repositories.NewUserRepo(dbPg)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userController := controllers.NewUserController(userUsecase)

	e.POST("user", userController.HHHHHHH)

}
