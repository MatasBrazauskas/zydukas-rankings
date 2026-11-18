package controller

import (
	"net/http"
	"zydukas-rankings/models"
	"zydukas-rankings/repo"
	"zydukas-rankings/service"
	"zydukas-rankings/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	userService *service.UserService
}

func CreateUserController(db *gorm.DB) *UserController {
	userRepo := repo.CreateUserRepo(db)
	userService := service.NewUserService(userRepo)

	return &UserController{
		userService: userService,
	}
}

func (u *UserController) CreateUser(g *gin.Context) {
	var user models.CreateUserPayload

	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": &utils.ErrInvalidUserData})
		return
	}

	userContext, err := g.Get(utils.UserContextTitle)
	if !err {
		g.JSON(http.StatusBadRequest, gin.H{"error": &utils.ErrBadUserCredentials})
		return
	}

	createUserContext := models.CreateUserContextPayload{Username: user.Username, FirebaseId: userContext.(models.UserContext).FireBaseId}

	if err := u.userService.CreateUser(&createUserContext); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	g.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}
