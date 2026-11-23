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

type PostController struct {
	postService *service.PostService
}

func CreatePostsController(db *gorm.DB, userRepo *repo.UserRepo) *PostController {
	postRepo := repo.CreatePostRepo(db)
	postService := service.NewPostService(postRepo, userRepo)

	return &PostController{
		postService: postService,
	}
}

func (p *PostController) CreatePost(g *gin.Context) {
	userContext, err := g.Get(utils.UserContextTitle)
	if !err {
		g.JSON(http.StatusBadRequest, gin.H{"error": &utils.ErrBadUserCredentials})
		return
	}

	firebaseId := userContext.(models.UserContext).FireBaseId

	caption := g.PostForm("caption")
	file, exists := g.FormFile("file")
	if exists != nil {
		g.JSON(http.StatusBadRequest, utils.ErrNoFileAttached)
		return
	}

	if err := p.postService.CreatePost(file, caption, firebaseId); err != nil {
		g.JSON(http.StatusBadRequest, err)
		return
	}

	g.JSON(http.StatusCreated, gin.H{"message": "File uploaded"})
}
