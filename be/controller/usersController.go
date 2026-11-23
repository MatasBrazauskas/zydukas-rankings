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
	UserService *service.UserService
}

func CreateUserController(db *gorm.DB) *UserController {
	userRepo := repo.CreateUserRepo(db)
	userService := service.NewUserService(userRepo)

	return &UserController{
		UserService: userService,
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

	if err := u.UserService.CreateUser(&createUserContext); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	g.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func (u *UserController) GetUser(g *gin.Context) {
	userContext, err := g.Get(utils.UserContextTitle)
	if !err {
		g.JSON(http.StatusBadRequest, gin.H{"error": &utils.ErrBadUserCredentials})
		return
	}

	firebaseId := userContext.(models.UserContext).FireBaseId

	user, uErr := u.UserService.GetUser(firebaseId)
	if uErr != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	g.JSON(http.StatusOK, gin.H{"data": user})
}

func (u *UserController) UpdateUsersUsername(g *gin.Context) {
	userContext, err := g.Get(utils.UserContextTitle)
	if !err {
		g.JSON(http.StatusBadRequest, gin.H{"error": &utils.ErrBadUserCredentials})
		return
	}

	firebaseId := userContext.(models.UserContext).FireBaseId
	newUsername := g.Param("new_username")
	if err := u.UserService.UpdateUsersUsername(firebaseId, newUsername); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "User updated successfully!"})
}

func (u *UserController) DeleteUser(g *gin.Context) {
	userContext, err := g.Get(utils.UserContextTitle)
	if !err {
		g.JSON(http.StatusBadRequest, gin.H{"error": &utils.ErrBadUserCredentials})
		return
	}

	firebaseId := userContext.(models.UserContext).FireBaseId
	if err := u.UserService.DeleteUser(firebaseId); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "User deleted successfully!"})
}
