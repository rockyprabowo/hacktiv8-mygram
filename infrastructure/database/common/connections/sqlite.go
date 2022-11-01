package database_connections

import (
	"errors"
	"github.com/spf13/viper"
	"rocky.my.id/git/mygram/infrastructure/configurations/config/keys"
)

func SQLiteConfig() (string, error) {
	databaseDSN := viper.GetString(config_keys.DatabaseDSN)
	if databaseDSN == "" {
		return "", errors.New("database.SQLiteConfig: DSN is empty")
	}
	return databaseDSN, nil
}
