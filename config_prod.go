//go:build prod && railway
// +build prod,railway

package main

import (
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/url"
	"rocky.my.id/git/mygram/configurations/config/keys"
)

func init() {
	// Set configuration file type and name.
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config.prod.railway")

	// Set the default
	viper.SetDefault(config_keys.Debug, false)
	viper.SetDefault(config_keys.Environment, "production")
	viper.SetDefault(config_keys.HostAddress, "")
	viper.SetDefault(config_keys.HostPort, "5002")
	viper.SetDefault(config_keys.BcryptCost, bcrypt.DefaultCost)

	// Bind to environment variables
	viper.MustBindEnv(config_keys.Environment, "RAILWAY_ENVIRONMENT")
	viper.MustBindEnv(config_keys.JWTSecret, "JWT_SECRET")
	viper.BindEnv(config_keys.HostPort, "PORT")
	viper.BindEnv(config_keys.BaseURL, "RAILWAY_STATIC_URL")
	setDatabaseDSN()

	if viper.GetString(config_keys.DatabaseEngine) == "" {
		dbURL, err := url.Parse(viper.GetString(config_keys.DatabaseDSN))
		if err != nil {
			log.Println("not parsing the database DSN as URL")
			return
		}
		viper.Set(config_keys.DatabaseEngine, dbURL.Scheme)
	}
}
