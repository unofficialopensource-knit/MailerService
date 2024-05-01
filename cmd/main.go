package main

import (
	"context"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"

	"github.com/unofficialopensource-knit/MailerService/internal/app"
)

func main() {
	appConfig, err := app.LoadConfig(context.Background())
	if err != nil {
		slog.Error(err.Error())
	}

	server := app.AppFactory(appConfig.Environment)

	if appConfig.LambdaTaskRoot != "" {
		fiberLambda := fiberadapter.New(server)
		lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
			return fiberLambda.ProxyWithContext(ctx, req)
		})
	} else {
		server.Listen(appConfig.BindAddress)
	}
}
