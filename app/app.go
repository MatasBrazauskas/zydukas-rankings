package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
	"zydukas-rankings/controller"
	"zydukas-rankings/middleware"
	"zydukas-rankings/utils"

	"firebase.google.com/go/auth"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	//"zydukas-rankings/controller"
)

type AppCore struct {
	UserController  *controller.UserController
	PostsController *controller.PostsController
}

func GetAppCore(db *sql.DB) *AppCore {
	return &AppCore{
		UserController:  controller.CreateUserController(db),
		PostsController: controller.CreatePostsController(db),
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

func GetHandlers(appCore *AppCore, authClient *auth.Client) http.Handler {
	router := gin.Default()

	router.GET("/health", controller.HealthCheckHandler) // public

	usersRoute := router.Group("/users")
	usersRoute.Use(middleware.FirebaseMiddleware(authClient))

	usersRoute.POST("/create", appCore.UserController.CreateUser)

	postsRoute := router.Group("/posts")
	postsRoute.Use(gzip.Gzip(gzip.DefaultCompression))

	postsRoute.POST("/upload", appCore.PostsController.CreatePost)

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
