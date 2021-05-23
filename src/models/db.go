// models/db.go

package models

import (
	"fmt"

	"github.com/josesalasdev/golang_api_template/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	mysqlDns := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MysqlUser,
		config.MysqlPassword,
		config.MysqlHost,
		config.MysqlPort,
		config.MysqlDataBase,
	)
	fmt.Println(mysqlDns)
	db, err := gorm.Open(mysql.Open(mysqlDns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Book{})

	DB = db
}
