package main

import (
	"github.com/spf13/viper"
	"log"
	"net/url"
	"rocky.my.id/git/mygram/configurations/config"
	"rocky.my.id/git/mygram/configurations/config/keys"
)

func setDatabaseEngine() {
	if !viper.IsSet(config_keys.DatabaseEngine) {
		dbURL, err := url.Parse(viper.GetString(config_keys.DatabaseDSN))
		if err != nil {
			log.Println("Couldn't parse the database DSN as URL. Assuming SQLite database engine.")
			viper.Set(config_keys.DatabaseEngine, "sqlite")
			return
		}
		if dbURL.Scheme == "file" {
			log.Println("Database URL starts with 'file'. Assuming SQLite database engine.")
			viper.Set(config_keys.DatabaseEngine, "sqlite")
			return
		}
		viper.Set(config_keys.DatabaseEngine, dbURL.Scheme)
	}
}

func setDatabaseDSN() {
	_ = viper.BindEnv(config_keys.DatabaseDSN, "DATABASE_URL")
	if !viper.IsSet(config_keys.DatabaseDSN) {
		_ = viper.BindEnv(config_keys.DatabaseDSN, "MYSQL_URL")
	}
	if !viper.IsSet(config_keys.DatabaseDSN) {
		log.Println("Database DSN is not set. Using SQLite in-memory database")
		viper.Set(config_keys.DatabaseEngine, config.DevDatabaseEngine)
		viper.Set(config_keys.DatabaseDSN, config.DevDatabaseDSN)
	}
}
