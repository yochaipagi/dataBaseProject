package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/yochaipagi/dataBaseProject/controller/api"
	"github.com/yochaipagi/dataBaseProject/controller/database"
)

func main() {

	err := godotenv.Load("/Users/yochaipagi/studying/Cs/dataBaseProject/controller/.env")
	if err != nil {
		log.Fatal("error loading .env file: " + err.Error())
	}

	err = database.SetupDB()
	if err != nil {
		log.Fatal("error setting up db: " + err.Error())
	}
	/*replicates := 10 // מספר השאילתות לביצוע לצורך המדידה
	dbSize := 2      // הכפלת בסיס הנתונים לצורך המדידה

	result, err := database.BenchmarkQuery(replicates, dbSize)
	if err != nil {
		log.Fatalf("Failed to benchmark the database: %v", err)
	}
	fmt.Printf("Benchmark results: %+v\n", result)
	*/ //if the api for benchmarking deosnt work proprly

	router := api.SetupRouter()
	err = router.Run(":9090")
	if err != nil {
		log.Fatal("error starting the server: " + err.Error())
	}

}
