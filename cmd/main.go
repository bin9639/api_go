package main

import (
	"api_go/cmd/api"
	"api_go/db"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewSQLDB(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Net:                  "tcp",
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatalf("error occurred while creating SQL DB: %v", err)
	}

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatalf("error occured while starting api server: %v", err)
	}
}
