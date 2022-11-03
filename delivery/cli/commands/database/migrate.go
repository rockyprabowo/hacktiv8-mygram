package database_commands

import (
	"context"
	"github.com/spf13/cobra"
	"log"
	"rocky.my.id/git/mygram/domain/entities"
	"rocky.my.id/git/mygram/infrastructure/database/common/connections"
)

var MigrateCmd = &cobra.Command{
	Use:   "db:migrate",
	Short: "Run the database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		MigrateDB(cmd.Context())
	},
}

func MigrateDB(ctx context.Context) {
	db := database_connections.Init()
	err := db.WithContext(ctx).
		AutoMigrate(
			&entities.User{},
			&entities.SocialMedia{},
			&entities.Photo{},
			&entities.Comment{},
		)
	if err != nil {
		log.Fatal("Migration failed, error: " + err.Error())
	}
}
