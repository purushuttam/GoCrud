package config

import (
	"fmt"
	"golang-fiber-crud/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDb(config *Config) *gorm.DB {
	sqlinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db, err := gorm.Open(postgres.Open(sqlinfo), &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("Db Connected Successfully")
	return db
}
