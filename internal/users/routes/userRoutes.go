package routes

import (
	"github.com/akmyrat/project1/internal/users/handler"
	"github.com/akmyrat/project1/internal/users/repository"
	"github.com/akmyrat/project1/internal/users/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func InitUserRoutes(router *gin.RouterGroup, DB *pgxpool.Pool) {
	userRepo := repository.NewUserRepository(DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	userRoutes := router.Group("/users")
	userRoutes.POST("/sign-up", userHandler.SignUp)
	userRoutes.POST("/sign-in", userHandler.SignIn)

}
