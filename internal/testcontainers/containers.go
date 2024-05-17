package testcontainers

import (
	"context"
	"fmt"
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"log/slog"
	"os"

	"github.com/fmiskovic/new-amz/internal/db"

	"time"

	"github.com/fmiskovic/new-amz/internal/utils"
	"github.com/fmiskovic/new-amz/migrations"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/migrate"
)

type TestDB struct {
	Ctx      context.Context
	BunDb    *bun.DB
	Shutdown func()
}

// SetUpDb is a helper func that runs postgres DB in a docker using testcontainers.
func SetUpDb() (*TestDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)

	// start postgres container
	postgres, err := startPostgresContainer(ctx)
	if err != nil {
		cancel()
		return nil, err
	}

	var bunDb *bun.DB
	for {
		if postgres.IsRunning() {
			slog.Info("postgres container is ready")
			host, err := postgres.Host(ctx)
			if err != nil {
				cancel()
				panic(err)
			}

			port, err := postgres.MappedPort(ctx, "5432")
			if err != nil {
				cancel()
				panic(err)
			}

			cfg := db.NewConfig().
				WithHost(fmt.Sprintf("%s:%d", host, port.Int())).
				WithDbName(utils.GetOrDefault("DB_NAME", "test-db")).
				WithUsername(utils.GetOrDefault("DB_USER", "test")).
				WithPassword(utils.GetOrDefault("DB_PASSWORD", "test")).
				Build()

			// connect db
			svc := db.NewService(db.WithConfig(*cfg))
			sqlDb, err := svc.Connect()
			if err != nil {
				cancel()
				return nil, err
			}

			bunDb = svc.WrapWithBun(sqlDb)

			// migrate db
			if err = migrateDB(ctx, bunDb); err != nil {
				cancel()
				return nil, err
			}

			// seed db
			bunDb.RegisterModel(
				(*entities.Entity)(nil),
				(*entities.Account)(nil),
				(*entities.Order)(nil),
				(*entities.Item)(nil),
				(*entities.OrderItem)(nil),
			)
			fixture := dbfixture.New(bunDb, dbfixture.WithTruncateTables())
			err = fixture.Load(ctx, os.DirFS("testdata"), "fixture.yml")
			if err != nil {
				cancel()
				return nil, err
			}
			break
		}
		slog.Info("waiting for postgres container...")
	}

	return &TestDB{
		Ctx:   ctx,
		BunDb: bunDb,
		Shutdown: func() {
			if err := terminateContainer(ctx, postgres); err != nil {
				slog.Warn("failed to terminate container", "warning", err)
			}
			cancel()
		},
	}, nil
}

func startPostgresContainer(ctx context.Context) (testcontainers.Container, error) {
	dbName := utils.GetOrDefault("DB_NAME", "test-db")
	dbUser := utils.GetOrDefault("DB_USER", "test")
	dbPassword := utils.GetOrDefault("DB_PASSWORD", "test")

	// Define a Postgres container configuration.
	req := testcontainers.ContainerRequest{
		Image:        "postgres:alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     dbUser,
			"POSTGRES_PASSWORD": dbPassword,
			"POSTGRES_DB":       dbName,
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections").WithOccurrence(2),
	}

	// create and start the postgres container.
	return testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		},
	)
}

func migrateDB(ctx context.Context, db *bun.DB) error {
	slog.Info("db migration is starting...")
	migrator := migrate.NewMigrator(db, migrations.Migrations)

	if err := migrator.Init(ctx); err != nil {
		return fmt.Errorf("init failed: %v", err)
	}

	if err := migrator.Lock(ctx); err != nil {
		slog.Warn("lock failed but it's ok, error message:", "warning", err)
	}
	defer migrator.Unlock(ctx) //nolint:errcheck

	group, err := migrator.Migrate(ctx)
	if err != nil {
		return fmt.Errorf("migration failed: %v", err)
	}

	if group.IsZero() {
		fmt.Printf("there are no new migrations to run (database is up to date)\n")
	}

	fmt.Printf("migrated to %s\n", group)

	return nil
}

func terminateContainer(ctx context.Context, container testcontainers.Container) error {
	if container == nil {
		slog.Info("container is nil, skipping terminate func")
		return nil
	}
	return container.Terminate(ctx)
}
