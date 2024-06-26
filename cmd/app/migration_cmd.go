package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/fmiskovic/new-amz/internal/db"
	"github.com/uptrace/bun"
	"log/slog"

	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
)

// newMigrationCmd configures set of migration cli commands.
func newMigrationCmd(migrations *migrate.Migrations) *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "database migrations",
		Subcommands: []*cli.Command{
			{
				Name:  "init",
				Usage: "create migration tables",
				Action: func(c *cli.Context) error {
					bunDb, err := connectDb()
					if err != nil {
						return err
					}
					migrator := migrate.NewMigrator(bunDb, migrations)
					return migrator.Init(c.Context)
				},
			},
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					bunDb, err := connectDb()
					if err != nil {
						return err
					}
					migrator := migrate.NewMigrator(bunDb, migrations)
					if err := migrator.Lock(c.Context); err != nil {
						return err
					}
					defer func(migrator *migrate.Migrator, ctx context.Context) {
						err := migrator.Unlock(ctx)
						if err != nil {
							slog.Error("failed to unlock", "error", err.Error())
						}
					}(migrator, c.Context)

					group, err := migrator.Migrate(c.Context)
					if err != nil {
						return err
					}
					if group.IsZero() {
						fmt.Printf("there are no new migrations to run (database is up to date)\n")
						return nil
					}
					fmt.Printf("migrated to %s\n", group)
					return nil
				},
			},
			{
				Name:  "rollback",
				Usage: "rollback the last migration group",
				Action: func(c *cli.Context) error {
					bunDb, err := connectDb()
					if err != nil {
						return err
					}
					migrator := migrate.NewMigrator(bunDb, migrations)
					if err := migrator.Lock(c.Context); err != nil {
						return err
					}
					defer func(migrator *migrate.Migrator, ctx context.Context) {
						err := migrator.Unlock(ctx)
						if err != nil {
							slog.Error("failed to unlock", "error", err.Error())
						}
					}(migrator, c.Context)

					group, err := migrator.Rollback(c.Context)
					if err != nil {
						return err
					}
					if group.IsZero() {
						fmt.Printf("there are no groups to roll back\n")
						return nil
					}
					fmt.Printf("rolled back %s\n", group)
					return nil
				},
			},
			{
				Name:  "lock",
				Usage: "lock migrations",
				Action: func(c *cli.Context) error {
					bunDb, err := connectDb()
					if err != nil {
						return err
					}
					migrator := migrate.NewMigrator(bunDb, migrations)
					return migrator.Lock(c.Context)
				},
			},
			{
				Name:  "unlock",
				Usage: "unlock migrations",
				Action: func(c *cli.Context) error {
					bunDb, err := connectDb()
					if err != nil {
						return err
					}
					migrator := migrate.NewMigrator(bunDb, migrations)
					return migrator.Unlock(c.Context)
				},
			},
			{
				Name:  "create_go",
				Usage: "create Go migration",
				Action: func(c *cli.Context) error {
					name := strings.Join(c.Args().Slice(), "_")
					bunDb, err := connectDb()
					if err != nil {
						return err
					}
					migrator := migrate.NewMigrator(bunDb, migrations)
					mf, err := migrator.CreateGoMigration(c.Context, name)
					if err != nil {
						return err
					}
					fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
					return nil
				},
			},
			{
				Name:  "create_sql",
				Usage: "create up and down SQL migrations",
				Action: func(c *cli.Context) error {
					name := strings.Join(c.Args().Slice(), "_")
					bunDb, err := connectDb()
					if err != nil {
						return err
					}
					migrator := migrate.NewMigrator(bunDb, migrations)
					files, err := migrator.CreateSQLMigrations(c.Context, name)
					if err != nil {
						return err
					}

					for _, mf := range files {
						fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
					}

					return nil
				},
			},
			{
				Name:  "status",
				Usage: "print migrations status",
				Action: func(c *cli.Context) error {
					bunDb, err := connectDb()
					if err != nil {
						return err
					}
					migrator := migrate.NewMigrator(bunDb, migrations)
					ms, err := migrator.MigrationsWithStatus(c.Context)
					if err != nil {
						return err
					}
					fmt.Printf("migrations: %s\n", ms)
					fmt.Printf("unapplied migrations: %s\n", ms.Unapplied())
					fmt.Printf("last migration group: %s\n", ms.LastGroup())
					return nil
				},
			},
			{
				Name:  "mark_applied",
				Usage: "mark migrations as applied without actually running them",
				Action: func(c *cli.Context) error {
					bunDb, err := connectDb()
					if err != nil {
						return err
					}
					migrator := migrate.NewMigrator(bunDb, migrations)
					group, err := migrator.Migrate(c.Context, migrate.WithNopMigration())
					if err != nil {
						return err
					}
					if group.IsZero() {
						fmt.Printf("there are no new migrations to mark as applied\n")
						return nil
					}
					fmt.Printf("marked as applied %s\n", group)
					return nil
				},
			},
		},
	}
}

func connectDb() (*bun.DB, error) {
	svc := db.NewService()
	sqlDb, err := svc.Connect()
	if err != nil {
		return nil, err
	}

	return svc.WrapWithBun(sqlDb), nil
}
