package cli

import (
	"log"
	"rocky.my.id/git/mygram/delivery/cli/commands"
	application "rocky.my.id/git/mygram/delivery/cli/commands/application"
	database "rocky.my.id/git/mygram/delivery/cli/commands/database"
)

func Execute() {
	commands.RootCmd.CompletionOptions.DisableDefaultCmd = true

	commands.RootCmd.AddCommand(application.DevCmd)
	commands.RootCmd.AddCommand(application.GenerateSecretsCmd)
	commands.RootCmd.AddCommand(application.ServeCmd)
	commands.RootCmd.AddCommand(database.MigrateCmd)
	commands.RootCmd.AddCommand(database.SeedCmd)

	if err := commands.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
