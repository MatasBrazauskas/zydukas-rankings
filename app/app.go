package main

import (
	"log"
	"net/http"
	"strconv"
	"zydukas-rankings/controller"
	"zydukas-rankings/middleware"
	"zydukas-rankings/utils"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	//"zydukas-rankings/controller"
)

type AppCore struct {
	logger          *zap.SugaredLogger
	UserController  *controller.UserController
	PostsController *controller.PostsController
}

func GetAppCore(db *gorm.DB, sugaredLogger *zap.SugaredLogger) *AppCore {
	return &AppCore{
		logger:          sugaredLogger,
		UserController:  controller.CreateUserController(db),
		PostsController: controller.CreatePostsController(db),
	}
}

type AppConfig struct {
	Port string
	/*writeTimeout time.Duration
	readTimeout  time.Duration
	idleTimeout  time.Duration*/
}

func GetAppConfig() *AppConfig {
	portNum, err := utils.EnvInt("SERVER_PORT")
	if err != nil {
		return nil
	}

	port := ":" + strconv.Itoa(portNum)

	return &AppConfig{
		Port: port,
		/*writeTimeout: utils.WriteTimeout,
		readTimeout:  utils.ReadTimeout,
		idleTimeout:  utils.IdleTimeout,*/
	}
}

func GetHandlers(appCore *AppCore, authClient *auth.Client) http.Handler {
	router := gin.Default()

	router.GET("/health", controller.HealthCheckHandler) // public

	usersRoute := router.Group("/users")
	usersRoute.Use(middleware.FirebaseMiddleware(authClient))

	usersRoute.POST("/create", appCore.UserController.CreateUser)

	postsRoute := router.Group("/posts")
	//postsRoute.Use(gzip.Gzip(gzip.DefaultCompression))

	postsRoute.POST("/upload", appCore.PostsController.CreatePost)

	return router
}

func RunApplication(handler http.Handler, appConfig *AppConfig) error {

	server := http.Server{
		Addr:    appConfig.Port,
		Handler: handler,
		/*WriteTimeout: appConfig.writeTimeout,
		ReadTimeout:  appConfig.readTimeout,
		IdleTimeout:  appConfig.idleTimeout,*/
	}

	log.Printf("Server started. Listening on %s", appConfig.Port)

	return server.ListenAndServe()
}
