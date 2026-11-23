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

func (u *UserRepo) GetUser(firebaseId string) (*models.User, error) {
	var user models.User

	if err := u.db.Where("firebase_id = ?", firebaseId).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) UpdateUsersUsername(user *models.User, newUsername string) error {
	return u.db.Model(user).Update("username", newUsername).Error
}

func (u *UserRepo) DeleteUser(firebaseId string) error {
	res := u.db.Where("firebase_id = ?", firebaseId).Delete(&models.User{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
