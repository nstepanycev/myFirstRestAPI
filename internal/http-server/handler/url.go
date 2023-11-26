package handler

import (
	"strconv"
	models "test/internal/models/save"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Create URL to database
func (h *Handler) CreateURL(c *gin.Context){
	var urls models.UrlToSaveService
	if err := c.BindJSON(&urls); err != nil{
		_ = c.Error(err)
		return
	}
	url, err := h.service.SaveURL(urls)
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
	url, err := h.service.URLSave.GetURL(id)
	if err != nil{
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, url)
}