package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

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
	// Yeni bir post olusturmak icin 'post' nesne hazirliyoruz
	var post model.Post

	// Http isteginden gelen json verisini 'post' nesnesine bagliyorus
	if err := c.ShouldBind(&post); err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Http istegindeki 'image' isimli dosyayi aliyoruz
	file, err := c.FormFile("image")
	if err != nil {
		handler.NewErrorResponse(c, 500, "invalid file")
		return
	}

	// Dosyanin uzantisini aliyoruz (jpg veya png)
	extension := filepath.Ext(file.Filename)

	// Yeni bir dosya adi olusturuyoruz
	newFileName := uuid.New().String() + extension

	// Resmi belirtilen dosya yoluna kaydetiyoruz
	if err := c.SaveUploadedFile(file, "uploads/images/"+newFileName); err != nil {
		handler.NewErrorResponse(c, 500, "failed storing image")
		return
	}

	// Resmin yolunu imgPath degiskenine atiyoruz ve post.ImagePath alanina bagliyorus
	ImgPath := "uploads/images/" + newFileName
	post.ImagePath = &ImgPath

	// Yeni postu veritabanina ekliyoruz
	newPost, err := h.service.CreatePost(&post)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Basarili olursa json formatinda yanit gonderilir
	c.JSON(200, gin.H{
		"data": newPost,
	})
}

func (h *PostHandler) DeletePostByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.DeletePostByID(id)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, "doesn't implement service")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}

func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.service.GetAllPosts()
	if err != nil {
		fmt.Print(err)
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"data": posts,
	})
}

func (h *PostHandler) GetPostByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, "invalid id")
		return
	}
	post, err := h.service.GetPostByID(id)
	if err != nil {
		fmt.Print(err)
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, post)
}

func (h *PostHandler) UpdatePostByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, "invalid id")
		return
	}

	var updatedPost model.Post
	if err := c.BindJSON(&updatedPost); err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.service.UpdatePostByID(id, &updatedPost); err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
	})
}
