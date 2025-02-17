package controller

import "github.com/gin-gonic/gin"

type PingController struct {
}

func (c *PingController) GetPing(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
