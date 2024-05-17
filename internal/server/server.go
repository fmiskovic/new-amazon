package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"

	"log"
	"log/slog"
)

// Server represents an HTTP server.
type Server struct {
	config Config
	router http.Handler
}

// Builder creates a new server instance.
type Builder struct {
	config Config
	router http.Handler
}

// NewBuilder creates a new Builder instance.
func NewBuilder() *Builder {
	return &Builder{}
}

// WithConfig sets the server configuration.
func (b *Builder) WithConfig(config Config) *Builder {
	b.config = config
	return b
}

func (b *Builder) WithRouter(r http.Handler) *Builder {
	b.router = r
	return b
}

func (b *Builder) Build() Server {
	if b.config.IsZero() {
		b.config = NewConfig().Build()
	}
	if b.router == nil {
		b.router = initRouter()
	}
	return Server{
		config: b.config,
		router: b.router,
	}
}

// Start starts the HTTP server and listens for incoming requests.
func (s Server) Start() {
	if s.router == nil {
		panic("server handler is not initialized")
	}

	server := http.Server{
		Addr:           s.config.addr,
		ReadTimeout:    s.config.readTimeout,
		WriteTimeout:   s.config.writeTimeout,
		MaxHeaderBytes: 1 << 20,
		Handler:        s.router,
	}

	// Start the server in a goroutine.
	go func() {
		slog.Info("Starting server", "address", s.config.addr)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
		slog.Info("Stopped serving new connections.")
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), s.config.shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	slog.Info("Graceful shutdown completed.")
}
