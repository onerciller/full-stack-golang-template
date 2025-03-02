package shared

import (
	"github.com/onerciller/fullstack-golang-template/pkg/httpserver"
	"github.com/onerciller/fullstack-golang-template/pkg/jwt"
)

type BaseHandler struct {
	JWT        jwt.Jwt
	HttpServer *httpserver.HttpServer
}

func NewBaseHandler(jwt jwt.Jwt, httpServer *httpserver.HttpServer) *BaseHandler {
	return &BaseHandler{
		JWT:        jwt,
		HttpServer: httpServer,
	}
}
