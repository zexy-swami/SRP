package main

import (
	"log"

	"github.com/zexy-swami/SRP/SRP_CLI/internal/SRP_net"
	"github.com/zexy-swami/SRP/SRP_CLI/pkg/parser"
)

const configName = "config"

func main() {
	if err := parser.ParseConfig(configName); err != nil {
		log.Fatalln(err.Error())
	}

	if err := SRP_net.StartClientConnection(); err != nil {
		log.Fatalln(err.Error())
	}
}
