package provider

import (
	"github.com/onerciller/fullstack-golang-template/internal/api"
	"github.com/onerciller/fullstack-golang-template/internal/api/base"
	"github.com/onerciller/fullstack-golang-template/internal/store"
	"github.com/onerciller/fullstack-golang-template/pkg/httpserver"
	"github.com/onerciller/fullstack-golang-template/pkg/jwt"
)

type Api struct{}

func (h *Api) Provide(
	userStore store.UserStore,
	jwtService jwt.Jwt,
	httpServer *httpserver.HttpServer,
) {
	api.Register(base.Handler{
		UserStore:  userStore,
		JwtService: jwtService,
		App:        httpServer.FiberApp,
	})
}
