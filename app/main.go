package main

import (
	"context"

	"github.com/mickaelyoshua/personal-finances/models"
)

func main() {
	c := context.Background()
	agent, err := models.NewAgent(c)
	if err != nil {
		panic("Failed to create SQL agent: " + err.Error())
	}
	defer agent.Conn.Close(c)
	server := models.NewServer(agent)

	err = server.Start("localhost:8080")
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}