package main

import (
	"context"
	"log"

	"github.com/mickaelyoshua/personal-finances/models"
	"github.com/mickaelyoshua/personal-finances/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load configuration: " + err.Error())
	}

	c := context.Background()
	agent, err := models.NewAgent(c, config.DatabaseURL)
	if err != nil {
		panic("Failed to create SQL agent: " + err.Error())
	}
	defer agent.Conn.Close(c)
	server := models.NewServer(agent)

	err = server.Start(config.ServerAddress)
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}