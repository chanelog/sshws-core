package main

import (
	"log"

	"github.com/chanelog/sshws-core/internal/config"
	"github.com/chanelog/sshws-core/internal/server"
)

func main() {

	cfg := config.Default()

	srv := server.New(cfg)

	log.Println("===================================")
	log.Println(" SSHWS Core")
	log.Println("===================================")

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
