package controller

import (
	"database/sql"
	"net/http"
	"os"
	"path/filepath"
	"zydukas-rankings/utils"

	"github.com/gin-gonic/gin"
)

type PostsController struct {
}

func CreatePostsController(db *sql.DB) *PostsController {
	return &PostsController{}
}

func (u *PostsController) CreatePost(g *gin.Context) {
	fh, err := g.FormFile("file")
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ErrNoFileAttached)
		return
	}

	const maxSize = 50 << 20 // 50 MB
	if fh.Size > maxSize {
		g.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "file too large"})
		return
	}

	if err := os.MkdirAll("output", 0o755); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create output folder"})
		return
	}

	path := filepath.Join("output", filepath.Base(fh.Filename))

	if err := g.SaveUploadedFile(fh, path); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "cannot save file"})
		return
	}

	g.JSON(http.StatusCreated, gin.H{"message": "File uploaded", "size": fh.Size})
}
