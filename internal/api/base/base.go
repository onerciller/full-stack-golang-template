package base

import (
	"github.com/gofiber/fiber/v2"
	"github.com/onerciller/fullstack-golang-template/internal/store"
	"github.com/onerciller/fullstack-golang-template/pkg/jwt"
)

type Handler struct {
	UserStore store.UserStore
	JWT       jwt.Jwt
	App       *fiber.App
}
