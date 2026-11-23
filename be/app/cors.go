package main

import (
	"time"
	"zydukas-rankings/utils"

	"github.com/gin-contrib/cors"
)

func CorsConfig() (*cors.Config, error) {
	allowOrigins, aErr := utils.EnvString("ALLOW_ORIGINS")
	if aErr != nil {
		return nil, aErr
	}

	corsConfig := cors.Config{
		AllowOrigins:     []string{allowOrigins},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	return &corsConfig, nil
}
