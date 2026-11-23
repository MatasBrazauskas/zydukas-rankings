package service

import (
	"crypto/rand"
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"
	"zydukas-rankings/repo"
	"zydukas-rankings/utils"

	"gorm.io/gorm"
)

type PostService struct {
	postRepo  *repo.PostRepo
	usersRepo *repo.UserRepo
}

func NewPostService(postRepo *repo.PostRepo, userRepo *repo.UserRepo) *PostService {
	return &PostService{
		postRepo:  postRepo,
		usersRepo: userRepo,
	}
}

func (p *PostService) CreatePost(file *multipart.FileHeader, caption, firebaseId string) *utils.CustomErrors {
	utils.Info("Firebase ID:" + firebaseId)

	user, err := p.usersRepo.GetUser(firebaseId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &utils.ErrUserNotFound
		}

		return &utils.ErrInvalidUserData
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	name := rand.Text() + ext
	path := filepath.Join("uploads", name)

	if err := p.postRepo.CreatePost(user.ID, caption, path, file); err != nil {
		return &utils.ErrCreatingPosts
	}

	return nil
}
