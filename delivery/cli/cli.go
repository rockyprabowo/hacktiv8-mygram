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
	database.SeedCmd.Flags().IntVarP(
		&database.DBSeedCount,
		"count",
		"c",
		20,
		"Count, affects user generation.",
	)
	database.SeedCmd.Flags().IntVarP(
		&database.DBSeedMultiplier,
		"multiplier",
		"m",
		5,
		"Count multiplier, affects social media, photo and comments generation.",
	)

	if err := commands.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
