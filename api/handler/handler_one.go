package handler

import "github.com/gin-gonic/gin"

type HandlerOneRequest struct {
	Request string `json:"request"`
}

func HandlerOne(ctx *gin.Context) {
	req := HandlerOneRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if req.Request != "success" {
		ctx.JSON(500, gin.H{
			"message": "Expected Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Success",
	})
}
