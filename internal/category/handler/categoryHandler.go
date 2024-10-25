package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/akmyrat/project1/internal/category/model"
	"github.com/akmyrat/project1/internal/category/service"
	handler "github.com/akmyrat/project1/pkg/response"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (h *CategoryHandler) PostCategory(c *gin.Context) {
	var category model.Category
	err := c.BindJSON(&category)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newCategory, err := h.service.CreateCategory(&category)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}
	c.JSON(200, gin.H{
		"data": newCategory,
	})
}

func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	Categories, err := h.service.GetAllCategories()
	if err != nil {
		c.JSON(500, "Something went wrong")
		return
	}

	c.JSON(200, gin.H{
		"data": Categories,
	})
}

func (h *CategoryHandler) GetOneCategory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid id")
	}

	category, err := h.service.GetOneCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}
	c.JSON(200, gin.H{
		"data": category,
	})
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	categoryID := c.Param("id")
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.DeleteCategoryByID(id)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, "does not get service")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category deleted successfully",
	})
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	categoryID := c.Param("id")
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var category model.Category
	if err := c.BindJSON(&category); err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.UpdateCategoryByiD(id, &category)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category updated successfully",
	})
}
