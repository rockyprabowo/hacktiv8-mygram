package application_commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"rocky.my.id/git/mygram/delivery/cli/commands/database"
	"rocky.my.id/git/mygram/infrastructure/configurations/config"
	"rocky.my.id/git/mygram/infrastructure/configurations/config/keys"
)

var DevCmd = &cobra.Command{
	Use:   "dev",
	Short: "Run the application with seeded in-memory SQLite database.",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set(config_keys.DatabaseEngine, config.DevDatabaseEngine)
		viper.Set(config_keys.DatabaseDSN, config.DevDatabaseDSN)

		database_commands.MigrateDB()
		database_commands.SeedDB(cmd.Context(), 100, 5)
		ServeApp()
	},
}