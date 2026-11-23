package repo

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"zydukas-rankings/models"

	"gorm.io/gorm"
)

type PostRepo struct {
	db *gorm.DB
}

func CreatePostRepo(db *gorm.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func saveFile(fh *multipart.FileHeader, path string) error {
	src, err := fh.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}

func (p *PostRepo) CreatePost(userId uint, caption, picturePath string, file *multipart.FileHeader) error {
	err := p.db.Transaction(func(tx *gorm.DB) error {
		post := models.Post{UserID: userId, Caption: caption, PicturePath: picturePath}
		if err := p.db.Create(&post).Error; err != nil {
			return err
		}

		if err := saveFile(file, post.PicturePath); err != nil {
			return err
		}

		return nil
	})

	return err
}
