package provider

import (
	"github.com/onerciller/fullstack-golang-template/internal/api"
	"github.com/onerciller/fullstack-golang-template/internal/entity"
	"github.com/onerciller/fullstack-golang-template/internal/store"
	"github.com/onerciller/fullstack-golang-template/pkg/config"
	"github.com/onerciller/fullstack-golang-template/pkg/database"
	"github.com/onerciller/fullstack-golang-template/pkg/httpserver"
	"github.com/onerciller/fullstack-golang-template/pkg/jwt"
	"github.com/samber/do"
	"gorm.io/gorm"
)

// Container manages dependency injection for the application
type Container struct {
	di *do.Injector
}

// NewContainer creates a new dependency injection container
func NewContainer() *Container {
	return &Container{di: do.New()}
}

// Register registers all application dependencies
func (c *Container) RegisterProviders() {
	do.Provide(c.di, config.Provide)
	do.Provide(c.di, httpserver.Provide)
	do.Provide(c.di, database.ProvidePostgres)
	do.Provide(c.di, jwt.ProvideJWTService)
	do.Provide(c.di, store.ProvideUserStore)
}

func (c *Container) RegisterServices() {
	// Get the http server
	httpServer := do.MustInvoke[*httpserver.HttpServer](c.di)

	// Get the user store
	userStore := do.MustInvoke[store.UserStore](c.di)

	// Get the jwt service
	jwtService := do.MustInvoke[jwt.Jwt](c.di)

	// Register the api
	api.Register(userStore, jwtService, httpServer.FiberApp)
}

func (c *Container) DatabaseMigrate() {
	db := do.MustInvoke[*gorm.DB](c.di)
	db.AutoMigrate(&entity.User{})
}

// Start starts the application
func (c *Container) Start() {
	httpServer := do.MustInvoke[*httpserver.HttpServer](c.di)
	httpServer.Start()
}

// Stop stops the application
func (c *Container) Stop() {
	httpServer := do.MustInvoke[*httpserver.HttpServer](c.di)
	httpServer.Shutdown()
}
