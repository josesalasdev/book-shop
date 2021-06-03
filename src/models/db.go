// models/db.go

package models

import (
	"fmt"

	"github.com/josesalasdev/golang_api_template/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Bogota",
		config.PostgresHostName,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDataBase,
		config.PostgresPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Category{})

	DB = db
}
