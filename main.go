package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Darkren/go-config/json"
	"github.com/Darkren/phonebook/server"
)

func main() {
	config, err := json.Load("appsettings.json")
	if err != nil {
		log.Fatalf("Got err loading config: %v", err)
	}

	server := &server.Server{}
	server.Start(config)
}
