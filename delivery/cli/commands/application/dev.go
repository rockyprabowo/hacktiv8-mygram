package application_commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"rocky.my.id/git/mygram/configurations/config"
	"rocky.my.id/git/mygram/configurations/config/keys"
	"rocky.my.id/git/mygram/delivery/cli/commands/database"
)

var DevCmd = &cobra.Command{
	Use:   "dev",
	Short: "Run the application with seeded in-memory SQLite database.",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set(config_keys.DatabaseEngine, config.DevDatabaseEngine)
		viper.Set(config_keys.DatabaseDSN, config.DevDatabaseDSN)

		database_commands.MigrateDB(cmd.Context())
		database_commands.SeedDBDev(cmd.Context())
		ServeApp()
	},
}
