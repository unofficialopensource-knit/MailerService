package main

import (
	"github.com/unofficialopensource-knit/MailerService/pkg/server"
)

func main() {
	conf := server.LoadConfig()

	router := server.AppFactory(conf.Environment)

	router.Run(conf.BindAddress)
}
