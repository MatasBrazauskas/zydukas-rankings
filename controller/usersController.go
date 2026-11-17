package controller

import (
	"database/sql"
	"net/http"
	"zydukas-rankings/models"
	"zydukas-rankings/repo"
	"zydukas-rankings/service"
	"zydukas-rankings/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func CreateUserController(db *sql.DB) *UserController {
	userRepo := repo.CreateUserRepo(db)
	userService := service.NewUserService(userRepo)

	return &UserController{
		userService: userService,
	}
}

func (u *UserController) CreateUser(g *gin.Context) {
	var user models.CreateUserPayload

	if err := g.ShouldBindJSON(&user); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": &utils.ErrCreatingUser})
		return
	}

	if err := u.userService.CreateUser(&user); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	g.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}
