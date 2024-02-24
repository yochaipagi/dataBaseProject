package main

import (
	"github.com/joho/godotenv"
	"github.com/ozhey/concordance/controller/api"
	"github.com/ozhey/concordance/controller/database"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file: " + err.Error())
	}

	err = database.SetupDB()
	if err != nil {
		log.Fatal("error setting up db: " + err.Error())
	}

	router := api.SetupRouter()
	err = router.Run(":9090")
	if err != nil {
		log.Fatal("error starting the server: " + err.Error())
	}
}
