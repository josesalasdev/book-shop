package config

import "os"

var Port = os.Getenv("PORT")
var MysqlDns = os.Getenv("MYSQL_DNS")
