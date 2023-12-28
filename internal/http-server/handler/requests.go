package handler

import (
	"net/http"
	"test/internal/api"
	models "test/internal/models/save"

	// "test/internal/service/storage"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


const aliasLength = 6


// Create URL to database
func (h *Handler) CreateURLService(c *gin.Context){

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
	url, err := h.service.CreateURLService(&req)
	if err != nil{
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, url)
}
// Get URL by id
func (h *Handler) GetURLbyIdService(c *gin.Context){
	var req models.Request
	if err := c.BindJSON(&req); err != nil{
		_ = c.Error(err)
		return
	}

	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil{
	// 	_ = c.Error(err)
	// 	return
	// }
	url, err := h.service.GetURLService(&req)
	if err != nil{
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, url)
}