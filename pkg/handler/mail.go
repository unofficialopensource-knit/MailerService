package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unofficialopensource-knit/MailerService/pkg/schema"
	"github.com/unofficialopensource-knit/MailerService/pkg/service"
)

func MailHandler(ctx *gin.Context) {
	var payload schema.MailRequestSchema
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.Error(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	go service.SendMail(payload)
	ctx.Status(http.StatusAccepted)
}
