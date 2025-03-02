package httpserver

import (
	"context"
	"fmt"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	_ "github.com/onerciller/fullstack-golang-template/docs" // This is required for Swagger documentation
	"github.com/schmentle/go-swagger-auth-form/swagger"
)

type Option func(*Config)

type Config struct {
	Port         string
	ServerHeader string
	AppName      string
	Version      string
	Middlewares  []fiber.Handler
}

type HttpServer struct {
	port     string
	FiberApp *fiber.App
}

func New(options ...Option) *HttpServer {
	config := &Config{}

	for _, option := range options {
		option(config)
	}

	instance := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler(),
		ServerHeader: config.ServerHeader,
		AppName:      config.AppName,
	})

	instance.Get("/swagger/doc.json", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})

	instance.Get("/swagger/*", adaptor.HTTPHandler(swagger.ServeSwaggerUI(swagger.SwaggerConfig{
		SwaggerDocURL: "/swagger/doc.json",
		AuthURL:       "/auth/login",
	})))

	for _, middleware := range config.Middlewares {
		instance.Use(middleware)
	}

	server := &HttpServer{
		FiberApp: instance,
		port:     config.Port,
	}

	return server
}

func (s *HttpServer) Shutdown() error {
	return s.FiberApp.Shutdown()
}

// ShutdownWithTimeout performs a graceful shutdown with a timeout context
func (s *HttpServer) ShutdownWithTimeout(ctx context.Context) error {
	// Create a channel for the shutdown result
	shutdownComplete := make(chan error, 1)

	// Start shutdown in a goroutine
	go func() {
		shutdownComplete <- s.FiberApp.Shutdown()
	}()

	// Wait for either shutdown completion or context timeout
	select {
	case err := <-shutdownComplete:
		return err
	case <-ctx.Done():
		// Forced shutdown if context is canceled or times out
		return ctx.Err()
	}
}

func (s *HttpServer) Start() error {
	addr := fmt.Sprintf(":%s", s.port)
	fmt.Printf("Server is starting on %s\n", addr)
	return s.FiberApp.Listen(addr)
}

func WithServerHeader(serverHeader string) Option {
	return func(s *Config) {
		s.ServerHeader = serverHeader
	}
}

func WithMiddlewares(middlewares ...fiber.Handler) Option {
	return func(s *Config) {
		s.Middlewares = middlewares
	}
}

func WithAppName(appName string) Option {
	return func(s *Config) {
		s.AppName = appName
	}
}

func WithPort(port string) Option {
	return func(s *Config) {
		s.Port = port
	}
}

func WithVersion(version string) Option {
	return func(s *Config) {
		s.Version = version
	}
}
