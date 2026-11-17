package main

import (
	"context"
	"log"
	"zydukas-rankings/utils"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if envErr := godotenv.Load(); envErr != nil {
		log.Fatal("Error loading .env file")
	}

	firebaseOptPath := utils.EnvString("FIREBASE_CONFIG_PATH", ".")
	opt := option.WithAuthCredentialsFile(option.ServiceAccount, firebaseOptPath)
	app, fErr := firebase.NewApp(context.Background(), nil, opt)

	if fErr != nil {
		log.Fatalf("error initializing app: %v", fErr)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatal("Failed to initialize firebase auth client.")
	}

	log.Print("Connected to Firebase API")

	dbConfig := utils.GetDatabaseConfig()

	db, dbErr := utils.InitDatabase(dbConfig)
	if dbErr != nil {
		log.Fatalf("Error initializing database: %v.", dbErr)
	}

	defer utils.CloseDatabase(db)

	appCore := GetAppCore(db)
	appConfig := GetAppConfig()

	handlers := GetHandlers(appCore, authClient)
	if runErr := RunApplication(handlers, appConfig); runErr != nil {
		log.Fatalf("Error running application: %v.", runErr)
	}
}
