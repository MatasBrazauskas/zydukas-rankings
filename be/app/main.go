package main

import (
	"fmt"
	"zydukas-rankings/middleware"
	"zydukas-rankings/utils"

	"github.com/bytedance/gopkg/util/logger"
	_ "github.com/lib/pq"
)

func main() {
	if err := utils.InitLogger(); err != nil {
		logger.Error(err.Error())
		return
	}

	if envErr := utils.InitEnv(); envErr != nil {
		utils.Fatal("Error loading .env file")
		return
	}

	utils.Info("logger works")

	authClient, err := middleware.InitFirebase()
	if err != nil {
		utils.Fatal("Failed to initialize firebase auth client.")
		return
	}
	utils.Info("Connected to Firebase API")

	db, dbErr := utils.InitDatabase()
	if dbErr != nil {
		msg := fmt.Sprintf("Failed to initialize database: %v", dbErr)
		utils.Fatal(msg)
		return
	}
	defer utils.CloseDatabase(db)
	utils.Info("Connected to database.")

	appCore := GetAppCore(db)
	appConfig := GetAppConfig()

	handlers := GetHandlers(appCore, authClient)
	if runErr := RunApplication(handlers, appConfig); runErr != nil {
		msg := fmt.Sprintf("Failed to initialize database: %v", dbErr)
		utils.Fatal(msg)
		utils.Fatal(msg)
	}
}
