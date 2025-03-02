package main

import (
	"context"
	"log"

	// Import the package
	_ "github.com/onerciller/fullstack-golang-template/docs" // This is required for Swagger documentation
	"github.com/onerciller/fullstack-golang-template/internal/auth"
	"github.com/onerciller/fullstack-golang-template/internal/shared"
	"github.com/onerciller/fullstack-golang-template/pkg/module"
	"gorm.io/gorm"
)

// @title Fullstack Golang Template API
// @version 1.0
// @description This is a sample server for a Fullstack Golang Template.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email your.email@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description "Enter your Bearer token in the format: `Bearer {token}`"
func main() {
	// Initialize the Fx application.
	app := module.RegisterWithDefault(
		&module.AppLifecycleModule{
			AutoMigrate: func(db *gorm.DB) error {
				return db.AutoMigrate(auth.UserEntity{})
			},
		},
		&shared.Module{},
		&auth.Module{},
	)

	// Start the application.
	if err := app.Start(context.Background()); err != nil {
		log.Fatal("Failed to start application:", err)
	}

	// Stop the application.
	if err := app.GracefulStop(context.Background()); err != nil {
		log.Fatal("Failed to stop application:", err)
	}
}
