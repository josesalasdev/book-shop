package config

import "os"

// Port of the app.
var Port = ":" + os.Getenv("PORT")

// SecretAPP of the app.
var SecretAPP = os.Getenv("SECRET")

// PostgresDataBase Postgres config.
var PostgresDataBase = os.Getenv("POSTGRES_DB")

// PostgresHostName Postgres config.
var PostgresHostName = os.Getenv("POSTGRES_HOSTNAME")

// PostgresUser Postgres config.
var PostgresUser = os.Getenv("POSTGRES_USER")

// PostgresPassword Postgres config.
var PostgresPassword = os.Getenv("POSTGRES_PASSWORD")

// PostgresPort Postgres config.
var PostgresPort = os.Getenv("POSTGRES_PORT")
