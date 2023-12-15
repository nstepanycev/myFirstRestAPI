package handler

import (
	"net/http"
	"strconv"
	"test/internal/api"
	models "test/internal/models/save"

	// "test/internal/service/storage"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


const aliasLength = 6

// Create URL to database
func (h *Handler) CreateURL(c *gin.Context){

	//JSON resp
	var req models.Request
	if err := c.BindJSON(&req); err != nil{
		_ = c.Error(err)
		return
	}
	// Validate
	v := validator.New()
	err := v.Struct(req)
	if err != nil{
		for _, e := range err.(validator.ValidationErrors){
			_ = c.Error(e)
			return
		}
	}
	//check aliace
	aliase := req.Aliase
	if aliase == ""{
		aliase = api.NewRandomString(aliasLength)
	}
	//Use Interface 
	url, err := h.service.SaveURL(&req)
	if err != nil{
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, url)
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