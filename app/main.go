package main

import (
	"context"
	"log"

	"github.com/mickaelyoshua/personal_finances/api"
	"github.com/mickaelyoshua/personal_finances/db/sqlc"
	"github.com/mickaelyoshua/personal_finances/util"
)

func main() {
	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal("Failed to load configuration: " + err.Error())
	}

	c := context.Background()
	agent, err := sqlc.NewAgent(c, config.DatabaseURL)
	if err != nil {
		panic("Failed to create SQL agent: " + err.Error())
	}

	//conn, err := util.GetConn(c, config.DatabaseURL)
	//if err != nil {
	//	log.Fatal("Failed to connect to database: " + err.Error())
	//}
	//defer conn.Close(c)
	//if err := util.ExecSQLScript(conn, "db/populate_categories.sql"); err != nil {
	//	log.Fatal("Failed to execute SQL script: " + err.Error())
	//}

	server := api.NewServer(agent)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Failed to start server: " + err.Error())
	}
}