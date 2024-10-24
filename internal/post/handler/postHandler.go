package handler

import (
	"net/http"
	"path/filepath"

	"github.com/akmyrat/project1/internal/post/model"
	"github.com/akmyrat/project1/internal/post/service"
	handler "github.com/akmyrat/project1/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var post model.Post
	err := c.ShouldBind(&post)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		handler.NewErrorResponse(c, 500, "invalid file")
		return
	}

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension

	if err := c.SaveUploadedFile(file, "uploads/images/"+newFileName); err != nil {
		handler.NewErrorResponse(c, 500, "failed storing image")
		return
	}
	imgPath := "uploads/images/" + newFileName
	post.ImagePath = &imgPath

	newPost, err := h.service.CreatePost(&post)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"data": newPost,
	})
}
