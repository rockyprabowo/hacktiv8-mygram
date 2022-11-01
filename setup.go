package main

import (
	"github.com/spf13/viper"
	"log"
	"rocky.my.id/git/mygram/infrastructure/configurations/config/keys"
	"strings"
)

func init() {
	// Set configuration file type and name.
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	// Set the default
	viper.SetDefault(config_keys.BaseURL, "")
	viper.SetDefault(config_keys.DatabaseEngine, "sqlite")
	viper.SetDefault(config_keys.DatabaseDSN, "file::memory:?cache=shared")
	viper.SetDefault(config_keys.Debug, true)
	viper.SetDefault(config_keys.JWTSecret, "!PLEASE-CHANGE!")
	viper.SetDefault(config_keys.Environment, "production")
	viper.SetDefault(config_keys.HostAddress, "")
	viper.SetDefault(config_keys.HostPort, "8005")
	viper.SetDefault(config_keys.BcryptCost, 4)

	// Bind to environment variables
	_ = viper.BindEnv(config_keys.HostPort, "PORT")
	_ = viper.BindEnv(config_keys.BaseURL, "RAILWAY_STATIC_URL")
	_ = viper.BindEnv(config_keys.Environment, "RAILWAY_ENVIRONMENT")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Read config
	if readErr := viper.ReadInConfig(); readErr != nil {
		//Write config file if it doesn't exist
		if _, ok := readErr.(viper.ConfigFileNotFoundError); ok {
			if writeErr := viper.SafeWriteConfig(); writeErr != nil {
				log.Fatalf("Failed to create config file(s): %s", writeErr)
			}
			return
		}
		log.Fatalf("Failed to load config file: %s", readErr)
	}
}
