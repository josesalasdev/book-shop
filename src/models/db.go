// models/db.go

package models

import (
	"github.com/josesalasdev/golang_api_template/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	db, err := gorm.Open(mysql.Open(config.MysqlDns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Book{})

	DB = db
}
