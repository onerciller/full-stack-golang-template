package httpserver

import (
	"fmt"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
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
		AuthURL:       "/auth/",
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

func (s *HttpServer) Start() error {
	return s.FiberApp.Listen(fmt.Sprintf(":%s", s.port))
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
