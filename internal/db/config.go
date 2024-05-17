package db

import (
	"log/slog"
	"runtime"
	"strconv"

	"github.com/fmiskovic/new-amz/internal/utils"
)

// Config represents the configuration for the database connection.
type Config struct {
	host        string // The host of the database.
	maxOpenConn int    // The maximum number of open connections.
	maxIdleConn int    // The maximum number of idle connections.
	dbName      string // The name of the database.
	username    string // The username for authentication.
	password    string // The password for authentication.
}

// ConfigBuilder is a builder for creating Config objects.
type ConfigBuilder struct {
	config *Config
}

// NewConfig creates a new ConfigBuilder instance.
func NewConfig() *ConfigBuilder {
	return &ConfigBuilder{
		config: &Config{},
	}
}

// WithUri sets the host of the database in the ConfigBuilder.
func (b *ConfigBuilder) WithHost(host string) *ConfigBuilder {
	b.config.host = host
	return b
}

// WithMaxOpenConn sets the maximum number of open connections in the ConfigBuilder.
func (b *ConfigBuilder) WithMaxOpenConn(maxOpenConn int) *ConfigBuilder {
	b.config.maxOpenConn = maxOpenConn
	return b
}

// WithMaxIdleConn sets the maximum number of idle connections in the ConfigBuilder.
func (b *ConfigBuilder) WithMaxIdleConn(maxIdleConn int) *ConfigBuilder {
	b.config.maxIdleConn = maxIdleConn
	return b
}

// WithDbName sets the name of the database in the ConfigBuilder.
func (b *ConfigBuilder) WithDbName(dbName string) *ConfigBuilder {
	b.config.dbName = dbName
	return b
}

// WithUsername sets the username for authentication in the ConfigBuilder.
func (b *ConfigBuilder) WithUsername(username string) *ConfigBuilder {
	b.config.username = username
	return b
}

// WithPassword sets the password for authentication in the ConfigBuilder.
func (b *ConfigBuilder) WithPassword(password string) *ConfigBuilder {
	b.config.password = password
	return b
}

// Build builds and returns the Config object.
func (b *ConfigBuilder) Build() *Config {
	if b.config.host == "" {
		b.config.host = utils.GetOrDefault("DB_HOST", "localhost:5432")
	}

	if b.config.maxOpenConn == 0 {
		numCpu := runtime.NumCPU() + 1
		maxOpenConn, err := strconv.Atoi(utils.GetOrDefault("DB_MAX_OPEN_CONN", strconv.Itoa(numCpu)))
		if err != nil {
			slog.Warn("error parsing DB_MAX_OPEN_CONN variable, using default", "error", err.Error())
			maxOpenConn = numCpu
		}
		b.config.maxOpenConn = maxOpenConn
	}

	if b.config.maxIdleConn == 0 {
		numCpu := runtime.NumCPU() + 1
		maxIdleConn, err := strconv.Atoi(utils.GetOrDefault("DB_MAX_IDLE_CONN", strconv.Itoa(numCpu)))
		if err != nil {
			slog.Warn("error parsing DB_MAX_IDLE_CONN variable, using default", "error", err.Error())
			maxIdleConn = numCpu
		}
		b.config.maxIdleConn = maxIdleConn
	}

	if b.config.dbName == "" {
		b.config.dbName = utils.GetOrDefault("DB_NAME", "go-db")
	}

	if b.config.username == "" {
		b.config.username = utils.GetOrDefault("DB_USER", "dbadmin")
	}

	if b.config.password == "" {
		b.config.password = utils.GetOrDefault("DB_PASSWORD", "dbadmin")
	}

	return b.config
}
