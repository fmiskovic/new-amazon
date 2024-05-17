package main

import (
	"log"
	"os"

	"github.com/fmiskovic/new-amz/internal/utils"
	"github.com/fmiskovic/new-amz/migrations"
	"github.com/urfave/cli/v2"
)

func main() {
	if err := utils.LoadEnv(); err != nil {
		panic(err)
	}

	app := &cli.App{
		Name: "app",

		Commands: []*cli.Command{
			newServeCmd(),
			newMigrationCmd(migrations.Migrations),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
