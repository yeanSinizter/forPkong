package repositories_test

import (
	"eiei/clean_architecture/db"
	"eiei/clean_architecture/entities"
	"eiei/clean_architecture/modules/user/repositories"
	"log"
	"testing"
)

func TestInsertData(t *testing.T) {
	dbPg, teardown := db.InitSql()

	defer teardown()

	userRepo := repositories.NewUserRepo(dbPg)

	user := &entities.UserModel{
		Username: "admin",
		Password: "admin",
	}

	userRes, err := userRepo.InsertUser(user)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("%#+v", userRes)

}
