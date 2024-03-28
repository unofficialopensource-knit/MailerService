package factory

import (
	"log"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/unofficialopensource-knit/MailerService/pkg/handler"
)

func App(mode string) *gin.Engine {
	switch strings.ToUpper(strings.TrimSpace(mode)) {
	case "DEBUG":
		gin.SetMode(gin.DebugMode)
	case "DEBUG-RELEASE":
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

	router.Use(cors.New(cors.Config{
		// Harden the cors origin
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/mail", handler.MailHandler)
	return router
}
