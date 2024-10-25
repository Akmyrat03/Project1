package routes

import (
	"github.com/akmyrat/project1/internal/post/handler"
	"github.com/akmyrat/project1/internal/post/repository"
	"github.com/akmyrat/project1/internal/post/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitPostRoutes(router *gin.RouterGroup, DB *sqlx.DB) {
	postRepo := repository.NewPostRepository(DB)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	postRoutes := router.Group("/post")
	postRoutes.POST("/add", postHandler.CreatePost)
}
