package main

import (
	"context"
	"log"
	"zydukas-rankings/utils"

	firebase "firebase.google.com/go"
	"go.uber.org/zap"
	"google.golang.org/api/option"

	_ "github.com/lib/pq"
)

func InitLogger() (*zap.SugaredLogger, error) {
	logger, sErr := zap.NewDevelopment()
	if sErr != nil {
		return nil, sErr
	}

	sugaredLogger := zap.New(logger.Core()).Sugar()
	return sugaredLogger, nil
}

func main() {
	if envErr := utils.InitEnv(); envErr != nil {
		log.Fatal("Error loading .env file")
		return
	}

	logger, lErr := InitLogger()
	if lErr != nil {
		log.Fatal("Error loading .env file")
		return
	}
	defer logger.Sync()

	firebaseOptPath, frErr := utils.EnvString("FIREBASE_CONFIG_PATH")
	if frErr != nil {
		logger.Fatal("Error reading FIREBASE_CONFIG_PATH")
		return
	}
	opt := option.WithAuthCredentialsFile(option.ServiceAccount, firebaseOptPath)
	app, fErr := firebase.NewApp(context.Background(), nil, opt)

	if fErr != nil {
		logger.Fatal("error initializing app: %v", fErr.Error())
		return
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		logger.Fatal("Failed to initialize firebase auth client.")
		return
	}

	logger.Info("Connected to Firebase API")

	db, dbErr := utils.InitDatabase()
	if dbErr != nil {
		logger.Fatal("Error initializing database: %v.", dbErr)
		return
	}
	defer utils.CloseDatabase(db)

	appCore := GetAppCore(db, logger)
	appConfig := GetAppConfig()

	handlers := GetHandlers(appCore, authClient)
	if runErr := RunApplication(handlers, appConfig); runErr != nil {
		logger.Fatal("Error running application: %v.", runErr)
	}
}
