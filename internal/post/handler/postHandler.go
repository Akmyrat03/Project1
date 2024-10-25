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
