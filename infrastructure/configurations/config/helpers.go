package config

import (
	"github.com/rockyprabowo/h8-helpers/slices"
	"github.com/spf13/viper"
	"rocky.my.id/git/mygram/infrastructure/configurations/config/keys"
)

// noinspection GoUnusedExportedFunction
func IsProduction(environment string) bool {
	return slices.StringInSlice(environment, ProductionEnvironments())
}

// noinspection GoUnusedExportedFunction
func IsDevelopment(environment string) bool {
	return slices.StringInSlice(environment, DevelopmentEnvironments())
}

func IsInProduction() bool {
	return slices.StringInSlice(viper.GetString(config_keys.Environment), ProductionEnvironments())
}

func IsInDevelopment() bool {
	return slices.StringInSlice(viper.GetString(config_keys.Environment), DevelopmentEnvironments())
}
