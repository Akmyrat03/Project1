package handler

import (
	"net/http"
	"strconv"

	"github.com/akmyrat/project1/internal/users/model"
	"github.com/akmyrat/project1/internal/users/service"
	handler "github.com/akmyrat/project1/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.service.CreateUser(input)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"username": input.Username,
	})
}

// gin is framework that is used for http requests
func (h *UserHandler) SignIn(c *gin.Context) {
	var input model.User
	//http request is sent json data, if we want store this data into variable we use c.BindJSON function
	if err := c.BindJSON(&input); err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.GenerateToken(input.Username, input.Password)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	if err := h.service.DeleteUser(id); err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
