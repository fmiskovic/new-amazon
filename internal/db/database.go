package db

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// Service represents a database service.
type Service struct {
	cfg Config
}

// ServiceOption is a function that configures a Service.
type ServiceOption func(*Service)

// NewService creates a new Service with the provided options or if options are not provider with default config.
func NewService(opts ...ServiceOption) *Service {
	s := &Service{}
	for _, opt := range opts {
		opt(s)
	}

	if opts == nil {
		s.cfg = *NewConfig().Build()
	}
	return s
}

// WithConfig sets the configuration for the Service.
func WithConfig(cfg Config) ServiceOption {
	return func(s *Service) {
		s.cfg = cfg
	}
}

// Connect connects to the database and returns a *sql.DB instance or an error if the connection fails.
func (svc Service) Connect() (*sql.DB, error) {
	conn := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", svc.cfg.username, svc.cfg.password, svc.cfg.host, svc.cfg.dbName)
	slog.Info("initializing db with conn string", "conn", conn)

	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(conn)))
	db.SetMaxOpenConns(svc.cfg.maxOpenConn)
	db.SetMaxIdleConns(svc.cfg.maxIdleConn)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// WrapWithBun wraps a *sql.DB instance with bun.DB.
// It can be useful for hooking bun.DB with bun.QueryHook.
func (svc Service) WrapWithBun(db *sql.DB) *bun.DB {
	bunDb := bun.NewDB(db, pgdialect.New())
	// if utils.IsDev() {
	// 	bunDb.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	// }
	return bunDb
}
