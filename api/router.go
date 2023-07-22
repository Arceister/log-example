package api

import (
	"log-example/api/handler"

	"github.com/gin-gonic/gin"
)

type Route struct {
	server *gin.Engine
}

func NewRoute(
	server *gin.Engine,
) *Route {
	return &Route{
		server: server,
	}
}

func (r *Route) Setup(g *gin.Engine) *gin.Engine {
	g.GET("/", handler.HandlerGet)

	g.POST("/", handler.HandlerOne)

	return g
}
