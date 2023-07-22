package handler

import "github.com/gin-gonic/gin"

func HandlerGet(ctx *gin.Context) {
	requestQuery := ctx.Query("request")
	if requestQuery == "error" {
		ctx.JSON(500, gin.H{
			"message": "Expected Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
	})
}
