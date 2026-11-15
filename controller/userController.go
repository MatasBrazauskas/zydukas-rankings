package controller

import (
	"database/sql"
	"log"
	"net/http"
	"zydukas-rankings/models"
	"zydukas-rankings/repo"
	"zydukas-rankings/service"
	"zydukas-rankings/utils"
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

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := utils.ReadJson(w, r, &user); err != nil {
		log.Println("Decode error:", err.Error())
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := u.userService.CreateUser(&user); err != nil {
		log.Println("Service error:", err.Error)
		http.Error(w, err.Error, http.StatusInternalServerError)
		return
	}

	if err := utils.WriteJson(w, http.StatusCreated, user); err != nil {
		log.Println("Encode error:", err.Error())
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
