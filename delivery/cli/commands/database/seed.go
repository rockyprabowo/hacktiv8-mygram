package database_commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"rocky.my.id/git/mygram/configurations/config"
	"rocky.my.id/git/mygram/infrastructure/database/comment"
	"rocky.my.id/git/mygram/infrastructure/database/common/connections"
	"rocky.my.id/git/mygram/infrastructure/database/photo"
	"rocky.my.id/git/mygram/infrastructure/database/social_media"
	"rocky.my.id/git/mygram/infrastructure/database/user"
)

var SeedCmd = &cobra.Command{
	Use:   "db:seed",
	Short: "Run the database seeder. Only works in development environment.",
	Run: func(cmd *cobra.Command, args []string) {
		var count, multiplier int

		cmd.Flags().IntVarP(
			&count,
			"count",
			"c",
			20,
			"Count, affects user generation.",
		)
		cmd.Flags().IntVarP(
			&multiplier,
			"multiplier",
			"m",
			5,
			"Count multiplier, affects social media, photo and comments generation.",
		)

		if config.IsInProduction() {
			fmt.Println("RUNNING IN PRODUCTION! Exiting...")
			return
		}

		SeedDB(cmd.Context(), count, multiplier)
	},
}

func SeedDB(ctx context.Context, count, multiplier int) {
	db := database_connections.Init()
	log.Println("Seeder started.")
	user_repository.Seed(ctx, db, count)
	social_media_repository.Seed(ctx, db, count, multiplier)
	photo_repository.Seed(ctx, db, count, multiplier)
	comment_repository.Seed(ctx, db, count, multiplier)
	log.Println("Seeder finished.")
}

func SeedDBDev(ctx context.Context) {
	SeedDB(ctx, 100, 5)
}
