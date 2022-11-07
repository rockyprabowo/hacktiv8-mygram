//go:build !prod

package main

import (
	"fmt"
	"github.com/spf13/viper"
	"rocky.my.id/git/mygram/configurations/config"
	"rocky.my.id/git/mygram/configurations/config/keys"
	"strings"
)

func init() {
	// Set configuration file type and name.

	// Set the default
	viper.SetDefault(config_keys.BaseURL, "localhost:8005")
	viper.SetDefault(config_keys.DatabaseEngine, "sqlite")
	viper.SetDefault(config_keys.DatabaseDSN, "file::memory:?cache=shared")
	viper.SetDefault(config_keys.Debug, true)
	viper.SetDefault(config_keys.JWTSecret, "!PLEASE-CHANGE!")
	viper.SetDefault(config_keys.Environment, "development")
	viper.SetDefault(config_keys.HostAddress, "localhost")
	viper.SetDefault(config_keys.HostPort, "8005")
	viper.SetDefault(config_keys.BcryptCost, 4)

	// Bind to environment variables
	_ = viper.BindEnv(config_keys.Environment, "RAILWAY_ENVIRONMENT")
	_ = viper.BindEnv(config_keys.HostPort, "PORT")
	_ = viper.BindEnv(config_keys.BaseURL, "RAILWAY_STATIC_URL")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Read config
	if config.IsInDevelopment() {
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.SetConfigName("config.dev")
		if readErr := viper.ReadInConfig(); readErr != nil {
			// Write config file if it doesn't exist
			if _, ok := readErr.(viper.ConfigFileNotFoundError); ok {
				if writeErr := viper.SafeWriteConfig(); writeErr != nil {
					fmt.Printf("Failed to create config file(s): %s\n", writeErr)
				}
			}
			fmt.Printf("Failed to load config file: %s . Using defaults\n", readErr)
		}
	}

	setDatabaseDSN()
	setDatabaseEngine()
}
