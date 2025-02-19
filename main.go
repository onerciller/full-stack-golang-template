package main

import (
	_ "github.com/onerciller/fullstack-golang-template/docs" // This is required for Swagger documentation
	"github.com/onerciller/fullstack-golang-template/provider"
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
	container := provider.NewContainer()
	container.RegisterProviders()
	container.RegisterServices()
	container.DatabaseMigrate()
	container.Start()
}
