package api

import (
	"github.com/onerciller/fullstack-golang-template/internal/api/auth"
	"github.com/onerciller/fullstack-golang-template/internal/api/base"
	"github.com/onerciller/fullstack-golang-template/internal/api/user"
	"github.com/onerciller/fullstack-golang-template/pkg/middleware"
)

func Register(bh base.Handler) {
	auth.New(bh).RegisterRoutes(bh.App.Group("/auth"))
	user.New(bh).RegisterRoutes(bh.App.Group("/api/v1", middleware.JWTAuth(bh.JWT)))
}
