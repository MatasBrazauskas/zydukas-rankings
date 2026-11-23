package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"
	"zydukas-rankings/controller"
	"zydukas-rankings/middleware"
	"zydukas-rankings/repo"
	"zydukas-rankings/service"
	"zydukas-rankings/utils"

	"firebase.google.com/go/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	//"zydukas-rankings/controller"
)

type AppCore struct {
	UserController  *controller.UserController
	PostsController *controller.PostController
}

func GetAppCore(db *gorm.DB) *AppCore {
	userRepo := repo.CreateUserRepo(db)
	userService := service.NewUserService(userRepo)
	userController := controller.UserController{UserService: userService}

	postsController := controller.CreatePostsController(db, userRepo)
	return &AppCore{
		UserController:  &userController,
		PostsController: postsController,
	}
}

type AppConfig struct {
	Port string
}

func GetAppConfig() *AppConfig {
	portNum, err := utils.EnvInt("SERVER_PORT")
	if err != nil {
		return nil
	}

	port := ":" + strconv.Itoa(portNum)

	return &AppConfig{
		Port: port,
	}
}

func GetHandlers(appCore *AppCore, authClient *auth.Client) http.Handler {
	corsConfig, err := CorsConfig()
	if err != nil {
		panic(err)
	}

	redisConfig, err := middleware.RedisConfig()
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(redisConfig)
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		zap.S().Fatalw("redis not reachable", "error", err)
	}

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20

	router.Use(cors.New(*corsConfig))
	router.Use(middleware.RedisMiddeware(rdb, 30*time.Second))

	router.GET("/health", controller.HealthCheckHandler)

	usersRoute := router.Group("/users")
	usersRoute.Use(middleware.FirebaseMiddleware(authClient))

	usersRoute.POST("", appCore.UserController.CreateUser)
	usersRoute.GET("", appCore.UserController.GetUser)
	usersRoute.PATCH(":new_username", appCore.UserController.UpdateUsersUsername)
	usersRoute.DELETE("", appCore.UserController.DeleteUser)

	postsRoute := router.Group("/posts")
	postsRoute.Use(middleware.FirebaseMiddleware(authClient))

	postsRoute.POST("/upload", appCore.PostsController.CreatePost)

	return router
}

func RunApplication(handler http.Handler, appConfig *AppConfig) error {

	server := http.Server{
		Addr:    appConfig.Port,
		Handler: handler,
	}

	log.Printf("Server started. Listening on %s", appConfig.Port)

	return server.ListenAndServe()
}
