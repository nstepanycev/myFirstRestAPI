package handler

import (
	"net/http"
	"strconv"
	models "test/internal/models/save"
	// "test/internal/service/storage"
	"github.com/gin-gonic/gin"
)

// Create URL to database
func (h *Handler) CreateURL(c *gin.Context){
	var urls *models.UrlToSaveService
	if err := c.BindJSON(&urls); err != nil{
		_ = c.Error(err)
		return
	}
	url, err := h.service.Storage.SaveURL(urls)
	if err != nil{
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated,url)
}
// Get URL by id
func (h *Handler) GetURLbyId(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		_ = c.Error(err)
		return
	}
	url, err := h.service.Storage.GetURL(id)
	if err != nil{
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, url)
}