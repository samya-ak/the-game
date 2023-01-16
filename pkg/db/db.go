package db

import (
	"fmt"
	"os"
	"thegame/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB_USER     = os.Getenv("POSTGRES_USER")
	DB_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	DB          = os.Getenv("POSTGRES_DB")
	DB_PORT     = os.Getenv("DB_PORT")
)

func Init() (*gorm.DB, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@database:%s/%s", DB_USER, DB_PASSWORD, DB_PORT, DB)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.GameState{})

	fmt.Println("Connected to DB ...")
	return db, nil
}
