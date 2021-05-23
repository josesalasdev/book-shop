package config

import "os"

var Port = os.Getenv("PORT")

// Mysql config
var MysqlDataBase = os.Getenv("MYSQL_DATABASE")
var MysqlUser = os.Getenv("MYSQL_USER")
var MysqlPassword = os.Getenv("MYSQL_PASSWORD")
var MysqlHost = os.Getenv("MYSQL_HOST")
var MysqlPort = os.Getenv("MYSQL_PORT")
