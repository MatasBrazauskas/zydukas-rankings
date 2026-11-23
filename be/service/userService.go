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

func (u *UserService) GetUser(firebaseId string) (*models.UserInfo, *utils.CustomErrors) {
	user, err := u.userRepo.GetUser(firebaseId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &utils.ErrUserNotFound
		}

		return nil, &utils.ErrInvalidUserData
	}

	userInfo := models.UserInfo{Username: user.Username, FirebaseId: user.FirebaseId, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}

	return &userInfo, nil
}

func (u *UserService) UpdateUsersUsername(firebaseId, newUsername string) *utils.CustomErrors {
	user, err := u.userRepo.GetUser(firebaseId)

	if err != nil {
		return &utils.ErrInvalidUserData
	}

	if err := u.userRepo.UpdateUsersUsername(user, newUsername); err != nil {
		return &utils.ErrUpdatingUsersUsername
	}

	return nil
}

func (u *UserService) DeleteUser(firebaseId string) *utils.CustomErrors {
	if err := u.userRepo.DeleteUser(firebaseId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &utils.ErrUserNotFound
		}

		return &utils.ErrInvalidUserData
	}

	return nil
}
