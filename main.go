package main

import (
	"basic-event-emitter/client"
	"basic-event-emitter/config"
	"basic-event-emitter/server"
	"log"
)

func main() {

	cfg := config.GetConfig()

	socketioIoClient, err := client.NewSocketClient(cfg)
	if err != nil {
		log.Println(err)
		return
	}

	srv := server.NewServer(cfg, socketioIoClient)

	srv.Serve()
}
