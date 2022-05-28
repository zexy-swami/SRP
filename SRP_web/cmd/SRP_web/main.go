package main

import (
	"log"

	"github.com/zexy-swami/SRP/SRP_web/internal/db"
	"github.com/zexy-swami/SRP/SRP_web/internal/router"
)

func main() {
	if err := db.PingDB(); err != nil {
		log.Fatalln(err.Error())
	}
	defer db.CloseDB()

	r := router.NewRouter()
	if err := r.RunTLS(":443", "cert/cert.crt", "cert/cert.key"); err != nil {
		log.Fatalln(err.Error())
	}
}
