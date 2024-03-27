package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"golang.org/x/net/context"

	"github.com/unofficialopensource-knit/MailerService/pkg/factory"
)

var ginLambda *ginadapter.GinLambda

func main() {
	conf, err := factory.Config()
	if err != nil {
		log.Panicf("Got error while loading config %v", err.Error())
	}

	router := factory.App(conf.Environment)

	if conf.Environment == "release" {
		ginLambda = ginadapter.New(router)
		lambda.Start(LambdaHandler)
	} else {
		fmt.Println("Low")
		router.Run(conf.BindAddress)
		fmt.Println("high")
	}
}

func LambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}
