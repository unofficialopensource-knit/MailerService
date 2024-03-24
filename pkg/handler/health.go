package handler

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/unofficialopensource-knit/MailerService/pkg/schema"
)

func HealthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, schema.HealthResponse{
		Message: "healthy",
		Version: runtime.Version(),
	})
}
