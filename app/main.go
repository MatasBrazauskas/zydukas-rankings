package main

import (
	"log"
	"zydukas-rankings/utils"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if envErr := godotenv.Load(); envErr != nil {
		log.Fatal("Error loading .env file")
	}

	dbConfig := utils.GetDatabaseConfig()

	db, dbErr := utils.InitDatabase(dbConfig)
	if dbErr != nil {
		log.Fatalf("Error initializing database: %v.", dbErr)
	}

	defer utils.CloseDatabase(db)

	appCore := GetAppCore(db)
	appConfig := GetAppConfig()

	handlers := GetHandlers(appCore)
	if runErr := RunApplication(handlers, appConfig); runErr != nil {
		log.Fatalf("Error running application: %v.", runErr)
	}
}
