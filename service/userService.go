package service

import (
	"zydukas-rankings/models"
	"zydukas-rankings/repo"
	"zydukas-rankings/utils"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService(userRepo *repo.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) CreateUser(user *models.CreateUserPayload) *utils.CustomErrors {
	if err := u.userRepo.CreateUser(user); err != nil {
		return &utils.ErrCreatingUser
	}

	return nil
}
