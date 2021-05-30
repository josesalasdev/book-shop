package config

import "os"

var Port = ":" + os.Getenv("PORT")

// Postgres config.
var PostgresDataBase = os.Getenv("POSTGRES_DB")
var PostgresHostName = os.Getenv("POSTGRES_HOSTNAME")
var PostgresUser = os.Getenv("POSTGRES_USER")
var PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
var PostgresPort = os.Getenv("POSTGRES_PORT")
