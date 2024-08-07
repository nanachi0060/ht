package main

import (
	"httpHandler/handler"
	"httpHandler/initialization"
	"httpHandler/libs"
	"log"
)

func main() {
	config, err := libs.GetConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	agents, err := initialization.InitExternalServicesAgents()
	if err != nil {
		log.Fatal(err)
	}

	err = handler.Run(config.ServicePort, agents, config)
	if err != nil {
		log.Fatal(err)
	}
}
