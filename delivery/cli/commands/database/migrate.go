package database_commands

import (
	"github.com/spf13/cobra"
	"log"
	"rocky.my.id/git/mygram/domain/entities"
	"rocky.my.id/git/mygram/infrastructure/database/common/connections"
)

var MigrateCmd = &cobra.Command{
	Use:   "db:migrate",
	Short: "Run the database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		MigrateDB()
	},
}

func MigrateDB() {
	db := database_connections.Init()
	err := db.AutoMigrate(&entities.User{}, &entities.SocialMedia{}, &entities.Photo{}, &entities.Comment{})
	if err != nil {
		log.Fatal("Migration failed, error: " + err.Error())
	}
}
