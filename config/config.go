package config

import (
	"github.com/joho/godotenv"
	"github.com/tkanos/gonfig"
	"log"
)

type Config struct {
	SocketServerHost string `env:"SOCKET_SERVER_HOST"`
	SockerServerPort int    `env:"SOCKET_SERVER_PORT"`
	Port             string `env:"PORT"`
}

func GetConfig() *Config {

	configuration := &Config{}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file")
	}

	err := gonfig.GetConf("", configuration)
	if err != nil {
		log.Fatal(err)
	}

	return configuration
}
