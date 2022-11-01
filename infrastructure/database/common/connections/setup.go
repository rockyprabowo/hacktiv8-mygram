package database_connections

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"rocky.my.id/git/mygram/infrastructure/configurations/config/keys"
	"strings"
)

// Init initialise the database connection.
func Init() (db *gorm.DB) {
	var (
		dialector      gorm.Dialector
		databaseEngine = viper.GetString(config_keys.DatabaseEngine)
	)

	switch strings.ToLower(databaseEngine) {
	case "postgresql":
		fallthrough
	case "postgres":
		dsn, err := PostgresConfig()
		if err != nil {
			log.Fatal("PostgreSQL configuration error: " + err.Error())
		}
		dialector = postgres.Open(dsn)
	case "sqlite":
		dsn, err := SQLiteConfig()
		if err != nil {
			log.Fatal("SQLite configuration error: " + err.Error())
		}
		dialector = sqlite.Open(dsn)
	case "mysql":
		fallthrough
	case "mariadb":
		dsn, err := MySQLConfig()
		if err != nil {
			log.Fatal("MySQL/MariaDB configuration error: " + err.Error())
		}
		dialector = mysql.Open(dsn)
	default:
		log.Fatal("Database engine is not supported: " + databaseEngine)
	}

	db, err := gorm.Open(dialector)
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	if viper.GetBool(config_keys.Debug) {
		db = db.Debug()
	}

	return
}
