package main

import (
	"context"
	"log"
	"zydukas-rankings/utils"

	firebase "firebase.google.com/go"
	"go.uber.org/zap"
	"google.golang.org/api/option"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if envErr := godotenv.Load(); envErr != nil {
		log.Fatal("Error loading .env file")
	}

	logger, sErr := zap.NewDevelopment()
	if sErr != nil {
		log.Fatal("Error initializing logger")
	}
	sugaredLogger := zap.New(logger.Core()).Sugar()
	defer sugaredLogger.Sync()

	firebaseOptPath := utils.EnvString("FIREBASE_CONFIG_PATH", ".")
	opt := option.WithAuthCredentialsFile(option.ServiceAccount, firebaseOptPath)
	app, fErr := firebase.NewApp(context.Background(), nil, opt)

	if fErr != nil {
		sugaredLogger.Fatal("error initializing app: %v", fErr.Error())

	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		logger.Fatal("Failed to initialize firebase auth client.")
	}

	logger.Info("Connected to Firebase API")

	dbConfig := utils.GetDatabaseConfig()

	db, dbErr := utils.InitDatabase(dbConfig)
	if dbErr != nil {
		sugaredLogger.Fatal("Error initializing database: %v.", dbErr)
	}

	defer utils.CloseDatabase(db)

	appCore := GetAppCore(db, sugaredLogger)
	appConfig := GetAppConfig()

	handlers := GetHandlers(appCore, authClient)
	if runErr := RunApplication(handlers, appConfig); runErr != nil {
		sugaredLogger.Fatal("Error running application: %v.", runErr)
	}
}
