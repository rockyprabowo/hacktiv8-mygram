package database_connections

import (
	"errors"
	"github.com/rockyprabowo/h8-helpers/options"
	"github.com/spf13/viper"
	"net"
	"net/url"
	"os"
	"rocky.my.id/git/mygram/configurations/config/keys"
	"strings"
)

// PostgresDBConfigMap is an alias of a map[string]string.
type PostgresDBConfigMap map[string]string

// PostgresConfig returns the database configurations.
func PostgresConfig() (string, error) {
	databaseDSN := viper.GetString(config_keys.DatabaseDSN)
	if databaseDSN != "" {
		config, err := parseDatabaseURL(databaseDSN)
		if err != nil {
			return "", err
		}
		return buildDSN(config)
	}
	config := parseFromEnvVars()
	return buildDSN(config)
}

// parseFromEnvVars returns a database configuration from loaded environment variables.
func parseFromEnvVars() PostgresDBConfigMap {
	return PostgresDBConfigMap{
		"host":     os.Getenv("DB_HOST"),
		"port":     os.Getenv("DB_PORT"),
		"user":     os.Getenv("DB_USER"),
		"password": os.Getenv("DB_PASSWORD"),
		"dbname":   os.Getenv("DB_DATABASE"),
		"sslmode":  os.Getenv("DB_SSLMODE"),
	}
}

// parseDatabaseURL returns a database configuration from the database URL defined in DB_URL environment variable.
func parseDatabaseURL(databaseURL string) (config PostgresDBConfigMap, err error) {
	var parsedURL *url.URL
	var host, port, username string

	parsedURL, err = url.Parse(databaseURL)
	if err != nil {
		return
	}

	host, port, err = net.SplitHostPort(parsedURL.Host)
	if err != nil {
		return
	}

	username = parsedURL.User.Username()
	password, hasPassword := parsedURL.User.Password()

	config = PostgresDBConfigMap{
		"host":    host,
		"port":    port,
		"user":    username,
		"dbname":  strings.Trim(parsedURL.Path, "/"),
		"sslmode": options.Default(os.Getenv("DB_SSLMODE"), "disable"),
	}

	if hasPassword {
		config["password"] = password
	}

	return
}

// buildDSN builds the DSN string from a database configuration.
func buildDSN(config PostgresDBConfigMap) (string, error) {
	var sb strings.Builder

	fields := []string{"host", "user", "dbname", "sslmode"}

	if _, hasPassword := config["password"]; hasPassword {
		fields = append(fields, "password")
	}

	for _, v := range fields {
		value, exist := config[v]

		if !exist || value == "" {
			return "", errors.New("Mandatory database configuration " + v + " is missing!")
		}
		sb.WriteString(v + "=" + value + " ")

	}

	return strings.Trim(sb.String(), " "), nil
}
