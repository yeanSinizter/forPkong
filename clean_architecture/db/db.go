package db

import (
	"eiei/clean_architecture/entities"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitSql() (*gorm.DB, func()) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	SqlDb, _ := db.DB()

	db.AutoMigrate(&entities.UserModel{})

	teardown := func() {
		defer SqlDb.Close()
	}

	return db, teardown
}
