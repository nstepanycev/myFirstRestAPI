package handler

import (
	"net/http"
	"strconv"
	"test/internal/api"
	models "test/internal/models/save"
	service "test/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const aliasLength = 6

type Handler struct{
	service *service.Service
}

func NewHandler(service *service.Service) *Handler{
	return &Handler{service: service}
}

func (h *Handler) ShortnerRoute(router *gin.Engine){
	api := router.Group("/v1/shortner")
	{
		api.POST("/",h.CreateURLService)
		api.GET("/:id",h.GetURLbyIdService)
	}
}

func (h *Handler) InitRouter() *gin.Engine{
	router := gin.New()
	h.ShortnerRoute(router)

	router.Run("0.0.0.0:8080")
	return router
	
}

// Create URL to database
func (h *Handler) CreateURLService(c *gin.Context){

	//JSON resp
	var req models.Request
	if err := c.BindJSON(&req); err != nil{
		_ = c.Error(http.ErrAbortHandler)
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

	urlCreate := models.Request{
		req.Id,
		req.URL,
		aliase,
	}
	
	//Create object in DB 
	resp, err :=  h.service.CreateURLService(&urlCreate)
	if err != nil{
		_ = c.Error(err)
		return
	}


	c.JSON(http.StatusCreated, resp)
}
// Get URL by id
func (h *Handler) GetURLbyIdService(c *gin.Context){

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		_ = c.Error(err)
	}
	url, err := h.service.GetURLService(id)
	if err != nil{
		_ = c.Error(err)
		return
	}
	
	c.JSON(http.StatusOK, url)
}