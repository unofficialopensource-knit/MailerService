package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

var routerLambda *ginadapter.GinLambda

type Config struct {
	Environment string `encvonfig:"environment"`
	BindAddress string `envconfig:"bind_addr"`
}

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

func HealthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "healthy"})
}

func main() {
	var config Config
	err := envconfig.Process("mailer", &config)
	if err != nil {
		log.Panicf("Got error while loading config %v", err.Error())
	}

	router := AppFactory(config.Environment)
	routerLambda = ginadapter.New(router)

	lambda.Start(LambdaHandler)
}

func LambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return routerLambda.ProxyWithContext(ctx, req)
}
