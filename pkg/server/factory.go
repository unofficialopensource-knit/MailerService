package server

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func AppFactory(mode string) *gin.Engine {
	switch strings.ToUpper(strings.TrimSpace(mode)) {
	case "DEBUG":
		gin.SetMode(gin.DebugMode)
	case "TEST":
		gin.SetMode(gin.TestMode)
	case "RELEASE":
		gin.SetMode(gin.ReleaseMode)
	default:
		panic("Mode not supported for running the server")
	}

	gin.DisableConsoleColor()

	router := gin.Default()
	router.GET("/health", HealthHandler)
	return router
}
