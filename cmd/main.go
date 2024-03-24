package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/unofficialopensource-knit/MailerService/pkg/server"
)

var routerLambda *ginadapter.GinLambda

func main() {
	var config server.Config
	err := envconfig.Process("mailer", &config)
	if err != nil {
		log.Panicf("Got error while loading config %v", err.Error())
	}

	router := server.AppFactory(config.Environment)
	routerLambda = ginadapter.New(router)

	lambda.Start(Handler)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginadapter.ProxyWithContext(ctx, req)
}
