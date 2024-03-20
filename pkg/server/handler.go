package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func MailHandler(ctx *gin.Context) {
	var payload MailRequestSchema

	err := ctx.BindJSON(&payload)
	if err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request schema"})
		return
	}

	go SendMail(payload)

	ctx.Status(http.StatusNoContent)
}
