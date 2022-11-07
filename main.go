package main

import (
	"rocky.my.id/git/mygram/delivery/cli"
)

// @title       MyGram
// @version     1.0
// @description This is a REST API for MyGram.
//
// @contact.name  API Support
// @contact.url   https://rocky.my.id/
// @contact.email rocky@lazycats.id

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	cli.Execute()
}
