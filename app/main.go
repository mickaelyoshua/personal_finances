package main

import (
	"context"
	"log"

	"github.com/mickaelyoshua/personal-finances/api"
	"github.com/mickaelyoshua/personal-finances/db/sqlc"
	"github.com/mickaelyoshua/personal-finances/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load configuration: " + err.Error())
	}

	c := context.Background()
	agent, err := sqlc.NewAgent(c, config.DatabaseURL)
	if err != nil {
		panic("Failed to create SQL agent: " + err.Error())
	}

	server := api.NewServer(agent)

	err = server.Start(config.ServerAddress)
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}