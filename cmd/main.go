package main

import (
	"log"

	"github.com/unofficialopensource-knit/MailerService/pkg/factory"
	"github.com/unofficialopensource-knit/MailerService/pkg/schema"
)

func main() {
	conf, err := schema.LoadConfig()
	if err != nil {
		log.Panicf("Got error while loading config %v", err.Error())
	}

	router := factory.App(conf.Environment)

	if conf.Environment == "debug" {
		router.Run(conf.BindAddress)
	}
}
