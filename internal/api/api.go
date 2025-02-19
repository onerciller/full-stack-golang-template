package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/onerciller/fullstack-golang-template/internal/api/user"
	"github.com/onerciller/fullstack-golang-template/internal/store"
	"github.com/onerciller/fullstack-golang-template/pkg/jwt"
)

// API version
const apiVersion = "v1"

// API struct holds the handlers.
// Dependencies are passed to the registration functions.
type API struct {
	userHandler *user.Handler
}

// New creates a new API instance and initializes handlers.
func New(userStore store.UserStore, jwtService jwt.Jwt) *API {
	userHandler := user.New(userStore, jwtService)
	return &API{userHandler: userHandler}
}

// RegisterPublicRoutes registers public routes.
func (a *API) RegisterPublicRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/register", a.userHandler.Register)
	auth.Post("/login", a.userHandler.Login)
}

// RegisterProtectedRoutes registers protected routes.
func (a *API) RegisterProtectedRoutes(router fiber.Router) {
	v1 := router.Group(fmt.Sprintf("/api/%s", apiVersion))
	v1.Get("/users", a.userHandler.GetUsers)
}

// RegisterAPIRoutes registers all API routes (both public and protected).
func (a *API) RegisterAPIRoutes(app *fiber.App) {
	a.RegisterPublicRoutes(app)
	a.RegisterProtectedRoutes(app)
}

func Register(
	userStore store.UserStore,
	jwtService jwt.Jwt,
	app *fiber.App,
) {
	api := New(userStore, jwtService)
	api.RegisterAPIRoutes(app)
}
