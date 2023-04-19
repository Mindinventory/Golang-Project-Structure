package database

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func TestDbConfigFromEnv() DbConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	var dbConfig DbConfig
	err = envconfig.Process("test_db", &dbConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	return dbConfig
}
