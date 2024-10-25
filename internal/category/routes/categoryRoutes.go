package routes

import (
	"github.com/akmyrat/project1/internal/category/handler"
	"github.com/akmyrat/project1/internal/category/repository"
	"github.com/akmyrat/project1/internal/category/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitCategoryRoutes(router *gin.RouterGroup, DB *sqlx.DB) {
	categoryRepo := repository.NewCategoryRepository(DB)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	categoryRoutes := router.Group("/category")
	categoryRoutes.POST("/post", categoryHandler.PostCategory)
	categoryRoutes.GET("/GetAll", categoryHandler.GetAllCategories)
	categoryRoutes.GET("/GetOne/:id", categoryHandler.GetOneCategory)
	categoryRoutes.DELETE("/delete/:id", categoryHandler.DeleteCategory)
	categoryRoutes.PUT("/update/:id", categoryHandler.UpdateCategory)

}
