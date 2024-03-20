package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"ping": "pong"})
}
