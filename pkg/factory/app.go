package factory

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/unofficialopensource-knit/MailerService/pkg/handler"
)

func App(mode string) *gin.Engine {
	switch strings.ToUpper(strings.TrimSpace(mode)) {
	case "DEBUG":
		gin.SetMode(gin.DebugMode)
	case "TEST":
		gin.SetMode(gin.TestMode)
	case "RELEASE":
		gin.SetMode(gin.ReleaseMode)
	default:
		log.Panicln("Mode not supported for running the server")
	}
	gin.DisableConsoleColor()

	router := gin.Default()

	router.GET("/health", handler.HealthHandler)

	return router
}
