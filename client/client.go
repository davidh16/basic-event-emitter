package client

import (
	"basic-event-emitter/config"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"log"
)

func NewSocketClient(cfg *config.Config) (*gosocketio.Client, error) {
	c, err := gosocketio.Dial(
		gosocketio.GetUrl(cfg.SocketServerHost, cfg.SockerServerPort, false),
		transport.GetDefaultWebsocketTransport(),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = c.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		log.Fatal("Disconnected")
	})
	if err != nil {
		log.Fatal(err)
	}

	return c, nil

}
