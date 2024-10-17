package main

import (
	"fmt"
	"log"
	"net"

	"server/config"
	"server/lobby"
)

func main() {
	config, error := config.LoadConfig("config.yml")

	if error != nil {
		log.Fatalln("Error when reading config.yml... \n", error.Error())

		return
	}

	listener, error := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port))

	if error != nil {
		log.Fatalln("Error when starting TCP server... \n", error.Error())

		return
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()

		lobby.HandleConnection(connection, err)
	}
}
