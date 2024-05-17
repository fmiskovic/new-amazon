package server

import (
	"time"

	"github.com/fmiskovic/new-amz/internal/utils"
)

// Config represents the server configuration.
type Config struct {
	addr            string        // addr is the server address.
	readTimeout     time.Duration // readTimeout is the maximum duration (seconds) for reading the entire request.
	writeTimeout    time.Duration // writeTimeout is the maximum duration (seconds) before timing out writes of the response.
	shutdownTimeout time.Duration // shutdownTimeout is the maximum duration (seconds) before timing out server shutdown.
	secret          string        // secret is being used to encrypt session store.
}

// ConfigBuilder is a builder for creating Config instances.
type ConfigBuilder struct {
	config *Config
}

// NewConfig creates a new ConfigBuilder instance.
func NewConfig() *ConfigBuilder {
	return &ConfigBuilder{config: &Config{}}
}

// WithAddr sets the server address.
func (b *ConfigBuilder) WithAddr(addr string) *ConfigBuilder {
	b.config.addr = addr
	return b
}

// WithReadTimeout sets the maximum duration for reading the entire request.
func (b *ConfigBuilder) WithReadTimeout(timeout time.Duration) *ConfigBuilder {
	b.config.readTimeout = timeout
	return b
}

// WithWriteTimeout sets the maximum duration before timing out writes of the response.
func (b *ConfigBuilder) WithWriteTimeout(timeout time.Duration) *ConfigBuilder {
	b.config.writeTimeout = timeout
	return b
}

func (b *ConfigBuilder) WithShutdownTimeout(timeout time.Duration) *ConfigBuilder {
	b.config.shutdownTimeout = timeout
	return b
}

func (b *ConfigBuilder) WithSecret(secret string) *ConfigBuilder {
	b.config.secret = secret
	return b
}

// Build creates a new Config instance based on the builder's configuration.
// If any configuration values are not set, default values will be used.
func (b *ConfigBuilder) Build() Config {
	if b.config.addr == "" {
		b.config.addr = utils.GetOrDefault("HTTP_LISTEN_ADDR", ":8080")
	}
	if b.config.readTimeout == 0 {
		timeout := utils.GetOrDefaultInt("HTTP_READ_TIMEOUT", 5)
		b.config.readTimeout = time.Duration(timeout) * time.Second
	}
	if b.config.writeTimeout == 0 {
		timeout := utils.GetOrDefaultInt("HTTP_WRITE_TIMEOUT", 10)
		b.config.writeTimeout = time.Duration(timeout) * time.Second
	}
	if b.config.shutdownTimeout == 0 {
		timeout := utils.GetOrDefaultInt("HTTP_SHUTDOWN_TIMEOUT", 10)
		b.config.shutdownTimeout = time.Duration(timeout) * time.Second
	}

	if b.config.secret == "" {
		b.config.secret = utils.GetOrDefault("AUTH_JWT_SECRET", "changeme")
	}
	return *b.config
}

func (c *Config) IsZero() bool {
	return c.addr == "" &&
		c.readTimeout == time.Duration(0) &&
		c.writeTimeout == time.Duration(0) &&
		c.shutdownTimeout == time.Duration(0) &&
		c.secret == ""
}
