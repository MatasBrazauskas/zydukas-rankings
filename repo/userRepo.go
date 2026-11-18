package repo

import (
	"zydukas-rankings/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func CreateUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) CreateUser(userContext *models.CreateUserContextPayload) error {
	user := models.User{Username: userContext.Username, FirebaseId: userContext.FirebaseId}

	if err := u.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
