package middleware

import (
	"log-example/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func UseLogger(logger *logger.Logger) gin.HandlerFunc {
	log := logger.Setup()

	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		rawQuery := ctx.Request.URL.RawQuery

		ctx.Next()

		latency := time.Since(start)

		clientIP := ctx.ClientIP()
		method := ctx.Request.Method
		statusCode := ctx.Writer.Status()

		error := ctx.Errors.ByType(gin.ErrorTypePrivate).String()

		if rawQuery != "" {
			path = path + "?" + rawQuery
		}

		event := log.Info()
		if error != "" {
			event = log.Error()
		}

		event.
			Int("status_code", statusCode).
			Str("method", method).
			Int("started_at", int(start.Unix())).
			Dur("latency", latency).
			Str("client", clientIP).
			Str("path", path).
			Msg(error)
	}
}
