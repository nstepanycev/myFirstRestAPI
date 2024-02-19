package handler

import (
	// "test/internal/http-server/middleware/logger"
	service "test/internal/services"
	// "test/internal/services/storage"

	"github.com/gin-gonic/gin"
	// "test/internal/models/save"
	// "github.com/go-playground/validator/v10"
	// "net/http"
)


type Handler struct{
	service *service.Service
}

func NewHandler(service *service.Service) *Handler{
	return &Handler{service: service}
}



func (h *Handler) ShortnerRoute(router *gin.Engine){
	api := router.Group("/v1/shortner")
	{
		api.GET("/:id",h.GetURLbyIdService)
		api.POST("/",h.CreateURLService)
	}
}

func (h *Handler) InitRouter() *gin.Engine{
	router := gin.New()
	router.Use(gin.Logger())
	h.ShortnerRoute(router)
	router.Run("localhost:8080")
	return router
	
}