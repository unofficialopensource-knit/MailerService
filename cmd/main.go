package main

import (
	"context"
	"log/slog"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"

	"github.com/unofficialopensource-knit/MailerService/internal/app"
)

func main() {
	appConfig, err := app.NewHTTPConfig(context.Background())
	if err != nil {
		slog.Error(err.Error())
	}

	server := app.NewAPIServer(appConfig)

	if appConfig.LambdaTaskRoot != "" {
		lambda.Start(httpadapter.New(server.Server.Handler).ProxyWithContext)
	} else {
		server.Server.ListenAndServe()
	}
}
