package service

import (
	"errors"
	"zydukas-rankings/models"
	"zydukas-rankings/repo"
	"zydukas-rankings/utils"

	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService(userRepo *repo.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) CreateUser(user *models.CreateUserContextPayload) *utils.CustomErrors {
	if err := u.userRepo.CreateUser(user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return &utils.ErrUserExists
		}

		return &utils.ErrCreatingUser
	}

	return nil
}
