package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
	"zydukas-rankings/controller"
	"zydukas-rankings/utils"

	//"zydukas-rankings/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type AppCore struct {
	UserController *controller.UserController
}

func GetAppCore(db *sql.DB) *AppCore {
	return &AppCore{
		UserController: controller.CreateUserController(db),
	}
}

type AppConfig struct {
	port         string
	writeTimeout time.Duration
	readTimeout  time.Duration
	idleTimeout  time.Duration
}

func GetAppConfig() *AppConfig {
	return &AppConfig{
		port:         utils.Port,
		writeTimeout: utils.WriteTimeout,
		readTimeout:  utils.ReadTimeout,
		idleTimeout:  utils.IdleTimeout,
	}
}

func GetHandlers(appCore *AppCore) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.ClientIPFromRemoteAddr)

	router.Get("/health", controller.HealthCheckHandler)
	router.Post("/create_user", appCore.UserController.CreateUser)

	return router
}

func RunApplication(handler http.Handler, appConfig *AppConfig) error {

	server := http.Server{
		Addr:         appConfig.port,
		Handler:      handler,
		WriteTimeout: appConfig.writeTimeout,
		ReadTimeout:  appConfig.readTimeout,
		IdleTimeout:  appConfig.idleTimeout,
	}

	log.Printf("Server started. Listening on %s", appConfig.port)

	return server.ListenAndServe()
}
