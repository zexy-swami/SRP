package main

import (
	"log"
	"net"

	"github.com/zexy-swami/SRP/SRP_CLI/internal/SRP_net"
	"github.com/zexy-swami/SRP/SRP_CLI/internal/db"
	"github.com/zexy-swami/SRP/SRP_CLI/pkg/parser"
)

const configName = "config_server"

func main() {
	if err := db.PingDB(); err != nil {
		log.Fatalln(err.Error())
	}
	defer db.CloseDB()

	if err := parser.ParseConfig(configName); err != nil {
		log.Fatalln(err.Error())
	}

	listener, err := net.Listen("tcp", getAddress())
	if err != nil {
		log.Fatalln(err.Error())
	}
	SRP_net.ListenerHandler(listener)
}

func getAddress() string {
	port := parser.GetDataFromConfig("port")
	return ":" + port
}
