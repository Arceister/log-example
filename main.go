package main

import (
	routes "log-example/api"
	"log-example/api/middleware"
	"log-example/logger"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := logger.NewLogger("debug", os.Stdout)
	server := gin.Default()
	server.Use(middleware.UseLogger(logger))
	route := routes.NewRoute(server)
	r := route.Setup(server)
	r.Run()
}
